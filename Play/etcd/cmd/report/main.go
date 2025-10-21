package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"graphqlscripts/internal/graphqlclient"
)

const defaultQuery = `
query Members($limit: Int!, $offset: Int!) {
  members(limit: $limit, offset: $offset) {
    id
    firstName
    lastName
    email
    phone
    plan
    country
    age
    signupDate
    monthlySpend
  }
}
`

type member struct {
	ID           string  `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Plan         string  `json:"plan"`
	Country      string  `json:"country"`
	Age          int     `json:"age"`
	SignupDate   string  `json:"signupDate"`
	MonthlySpend float64 `json:"monthlySpend"`
}

type membersResponse struct {
	Members []member `json:"members"`
}

type reportSummary struct {
	GeneratedAt      time.Time         `json:"generatedAt"`
	TotalMembers     int               `json:"totalMembers"`
	AverageAge       float64           `json:"averageAge"`
	AverageSpend     float64           `json:"averageMonthlySpend"`
	PlanDistribution map[string]int    `json:"planDistribution"`
	CountrySpread    map[string]int    `json:"countryDistribution"`
	SignupEarliest   string            `json:"signupEarliest"`
	SignupLatest     string            `json:"signupLatest"`
	TopSpenders      []topSpenderEntry `json:"topSpenders"`
}

type topSpenderEntry struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Plan         string  `json:"plan"`
	MonthlySpend float64 `json:"monthlySpend"`
	Country      string  `json:"country"`
}

func main() {
	var (
		flagEndpoint  = flag.String("endpoint", envDefault("GRAPHQL_ENDPOINT", ""), "GraphQL endpoint URL")
		flagAuthToken = flag.String("auth-token", envDefault("GRAPHQL_BEARER_TOKEN", ""), "Bearer token for the Authorization header")
		flagQueryFile = flag.String("query-file", "", "Optional path to a file containing the GraphQL query to use")
		flagLimit     = flag.Int("limit", envInt("REPORT_LIMIT", 0), "Maximum number of members to fetch (0 means fetch everything available)")
		flagPageSize  = flag.Int("page-size", envInt("REPORT_PAGE_SIZE", 100), "Page size used for pagination")
		flagTimeout   = flag.Duration("timeout", 10*time.Second, "HTTP timeout for GraphQL requests")
		flagJSON      = flag.Bool("json", false, "Emit the report as JSON instead of human readable text")
		topCount      = flag.Int("top", envInt("REPORT_TOP", 5), "Number of top spenders to highlight")
	)
	flag.Parse()

	if *flagEndpoint == "" {
		exitWithError("missing GraphQL endpoint (use -endpoint or GRAPHQL_ENDPOINT)")
	}

	query := strings.TrimSpace(defaultQuery)
	if file := strings.TrimSpace(*flagQueryFile); file != "" {
		content, err := os.ReadFile(file)
		if err != nil {
			exitWithError("read query file: %v", err)
		}
		query = string(content)
	}

	if *flagPageSize <= 0 {
		exitWithError("page size must be greater than zero")
	}

	if *topCount <= 0 {
		exitWithError("top count must be greater than zero")
	}

	ctx := context.Background()
	client := graphqlclient.NewClient(*flagEndpoint, *flagAuthToken, *flagTimeout)

	results := collectMembers(ctx, client, query, *flagPageSize, *flagLimit)
	if len(results) == 0 {
		fmt.Println("No members returned from the GraphQL API")
		return
	}

	summary := buildReport(results, *topCount)
	if *flagJSON {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(summary); err != nil {
			exitWithError("encode report json: %v", err)
		}
		return
	}

	fmt.Printf("Report generated at %s\n", summary.GeneratedAt.Format(time.RFC3339))
	fmt.Printf("Members fetched: %d\n", summary.TotalMembers)
	fmt.Printf("Average age: %.1f years\n", summary.AverageAge)
	fmt.Printf("Average monthly spend: $%.2f\n", summary.AverageSpend)
	fmt.Println()

	fmt.Println("Plan distribution:")
	printSortedMap(summary.PlanDistribution)
	fmt.Println()

	fmt.Println("Country distribution:")
	printSortedMap(summary.CountrySpread)
	fmt.Println()

	fmt.Printf("Signup window: %s -> %s\n", summary.SignupEarliest, summary.SignupLatest)
	fmt.Println()

	fmt.Printf("Top %d spenders:\n", len(summary.TopSpenders))
	for i, entry := range summary.TopSpenders {
		fmt.Printf("  %d. %s (%s) plan=%s country=%s spend=$%.2f\n", i+1, entry.Name, entry.Email, entry.Plan, entry.Country, entry.MonthlySpend)
	}
}

func collectMembers(ctx context.Context, client *graphqlclient.Client, query string, pageSize, limit int) []member {
	var (
		offset  = 0
		results []member
	)

	for {
		var resp membersResponse
		vars := map[string]interface{}{
			"limit":  pageSize,
			"offset": offset,
		}

		if err := client.Execute(ctx, query, vars, &resp); err != nil {
			exitWithError("graphql query failed: %v", err)
		}

		if len(resp.Members) == 0 {
			break
		}

		results = append(results, resp.Members...)
		offset += len(resp.Members)

		if limit > 0 && len(results) >= limit {
			results = results[:limit]
			break
		}

		if len(resp.Members) < pageSize {
			break
		}
	}

	return results
}

func buildReport(members []member, topCount int) reportSummary {
	var (
		totalSpend float64
		totalAge   int
		validSpend int
		validAge   int
		firstSeen  time.Time
		lastSeen   time.Time
	)

	plans := map[string]int{}
	countries := map[string]int{}
	for _, m := range members {
		plans[m.Plan]++
		countries[m.Country]++

		if m.MonthlySpend > 0 {
			totalSpend += m.MonthlySpend
			validSpend++
		}

		if m.Age > 0 {
			totalAge += m.Age
			validAge++
		}

		if parsed, err := time.Parse(time.RFC3339, m.SignupDate); err == nil {
			if firstSeen.IsZero() || parsed.Before(firstSeen) {
				firstSeen = parsed
			}
			if parsed.After(lastSeen) {
				lastSeen = parsed
			}
		}
	}

	top := make([]topSpenderEntry, 0, len(members))
	for _, m := range members {
		top = append(top, topSpenderEntry{
			ID:           m.ID,
			Name:         strings.TrimSpace(m.FirstName + " " + m.LastName),
			Email:        m.Email,
			Plan:         m.Plan,
			MonthlySpend: m.MonthlySpend,
			Country:      m.Country,
		})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].MonthlySpend > top[j].MonthlySpend
	})
	if len(top) > topCount {
		top = top[:topCount]
	}

	var avgSpend float64
	if validSpend > 0 {
		avgSpend = totalSpend / float64(validSpend)
	}

	var avgAge float64
	if validAge > 0 {
		avgAge = float64(totalAge) / float64(validAge)
	}

	summary := reportSummary{
		GeneratedAt:      time.Now().UTC(),
		TotalMembers:     len(members),
		AverageAge:       round(avgAge, 1),
		AverageSpend:     round(avgSpend, 2),
		PlanDistribution: plans,
		CountrySpread:    countries,
		TopSpenders:      top,
	}

	if !firstSeen.IsZero() {
		summary.SignupEarliest = firstSeen.Format(time.RFC3339)
	}
	if !lastSeen.IsZero() {
		summary.SignupLatest = lastSeen.Format(time.RFC3339)
	}

	return summary
}

func printSortedMap(values map[string]int) {
	type kv struct {
		Key   string
		Value int
	}
	all := make([]kv, 0, len(values))
	for k, v := range values {
		all = append(all, kv{Key: k, Value: v})
	}

	sort.Slice(all, func(i, j int) bool {
		if all[i].Value == all[j].Value {
			return all[i].Key < all[j].Key
		}
		return all[i].Value > all[j].Value
	})

	for _, item := range all {
		fmt.Printf("  %s: %d\n", item.Key, item.Value)
	}
}

func round(val float64, precision int) float64 {
	if precision <= 0 {
		return float64(int(val + 0.5))
	}
	scale := mathPow10(precision)
	return mathRound(val*scale) / scale
}

func mathRound(v float64) float64 {
	if v < 0 {
		return float64(int(v - 0.5))
	}
	return float64(int(v + 0.5))
}

func mathPow10(exp int) float64 {
	value := 1.0
	for i := 0; i < exp; i++ {
		value *= 10
	}
	return value
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
