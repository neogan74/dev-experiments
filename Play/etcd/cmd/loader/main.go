package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"graphqlscripts/internal/graphqlclient"
)

const defaultMutation = `
mutation CreateMember($input: MemberInput!) {
  createMember(input: $input) {
    id
  }
}
`

type memberInput struct {
	ID           string  `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Age          int     `json:"age"`
	Country      string  `json:"country"`
	Plan         string  `json:"plan"`
	SignupDate   string  `json:"signupDate"`
	MonthlySpend float64 `json:"monthlySpend"`
}

type createMemberResponse struct {
	CreateMember struct {
		ID string `json:"id"`
	} `json:"createMember"`
}

func main() {
	var (
		flagEndpoint     = flag.String("endpoint", envDefault("GRAPHQL_ENDPOINT", ""), "GraphQL endpoint URL")
		flagAuthToken    = flag.String("auth-token", envDefault("GRAPHQL_BEARER_TOKEN", ""), "Bearer token used for the Authorization header")
		flagCount        = flag.Int("count", envInt("MEMBER_COUNT", 25), "Number of members to generate")
		flagMutationFile = flag.String("mutation-file", "", "Path to a file containing the GraphQL mutation to use")
		flagTimeout      = flag.Duration("timeout", 10*time.Second, "HTTP timeout for GraphQL requests")
		flagDryRun       = flag.Bool("dry-run", false, "Only print the generated payloads without sending them")
	)
	flag.Parse()

	if *flagEndpoint == "" {
		exitWithError("missing GraphQL endpoint (use -endpoint or GRAPHQL_ENDPOINT)")
	}

	mutation := strings.TrimSpace(defaultMutation)
	if file := strings.TrimSpace(*flagMutationFile); file != "" {
		content, err := os.ReadFile(file)
		if err != nil {
			exitWithError("read mutation file: %v", err)
		}
		mutation = string(content)
	}

	if *flagCount <= 0 {
		exitWithError("count must be greater than zero")
	}

	gofakeit.Seed(time.Now().UnixNano())

	ctx := context.Background()
	client := graphqlclient.NewClient(*flagEndpoint, *flagAuthToken, *flagTimeout)
	plans := []string{"FREE", "STANDARD", "PREMIUM"}
	countries := []string{"US", "CA", "GB", "DE", "FR", "ES", "NL", "IN", "BR", "AU"}

	for i := 0; i < *flagCount; i++ {
		plan := gofakeit.RandomString(plans)
		spend := 0.0
		if plan != "FREE" {
			spend = roundCurrency(gofakeit.Price(10, 250))
		}

		input := memberInput{
			ID:           gofakeit.UUID(),
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			Email:        gofakeit.Email(),
			Phone:        gofakeit.Phone(),
			Age:          gofakeit.Number(18, 75),
			Country:      gofakeit.RandomString(countries),
			Plan:         plan,
			SignupDate:   gofakeit.DateRange(time.Now().AddDate(-2, 0, 0), time.Now()).Format(time.RFC3339),
			MonthlySpend: spend,
		}

		payload := map[string]interface{}{
			"input": input,
		}

		fmt.Printf("member[%d]: %s %s (%s) plan=%s spend=%.2f country=%s\n",
			i+1, input.FirstName, input.LastName, input.Email, input.Plan, input.MonthlySpend, input.Country)

		if *flagDryRun {
			continue
		}

		var resp createMemberResponse

		if err := client.Execute(ctx, mutation, payload, &resp); err != nil {
			exitWithError("graphql mutation failed: %v", err)
		}

		fmt.Printf("  -> created id=%s\n", resp.CreateMember.ID)
	}
}

func roundCurrency(v float64) float64 {
	return math.Round(v*100) / 100
}

func envDefault(key, fallback string) string {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		return val
	}
	return fallback
}

func envInt(key string, fallback int) int {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return fallback
	}
	var parsed int
	if _, err := fmt.Sscanf(val, "%d", &parsed); err == nil {
		return parsed
	}
	return fallback
}

func exitWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
