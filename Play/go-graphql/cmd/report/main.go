package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/machinebox/graphql"
)

type reportResponse struct {
	Todos []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
		Done bool   `json:"done"`
		User struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"user"`
	} `json:"todos"`
}

func main() {
	endpoint := flag.String("endpoint", "http://localhost:8080/query", "GraphQL endpoint to fetch data from")
	flag.Parse()

	client := graphql.NewClient(*endpoint)
	req := graphql.NewRequest(`
		query ReportTodos {
			todos {
				id
				text
				done
				user {
					id
					name
				}
			}
		}
	`)

	var resp reportResponse
	if err := client.Run(context.Background(), req, &resp); err != nil {
		log.Fatalf("fetch todos: %v", err)
	}

	if len(resp.Todos) == 0 {
		fmt.Println("No todos returned from the API.")
		return
	}

	total := len(resp.Todos)
	doneCount := 0

	type userStat struct {
		name    string
		total   int
		done    int
		pending int
	}

	userStats := map[string]*userStat{}

	for _, todo := range resp.Todos {
		if todo.Done {
			doneCount++
		}

		stat, ok := userStats[todo.User.ID]
		if !ok {
			stat = &userStat{
				name: todo.User.Name,
			}
			userStats[todo.User.ID] = stat
		}

		stat.total++
		if todo.Done {
			stat.done++
		} else {
			stat.pending++
		}
	}

	fmt.Printf("Fetched %d todos (done: %d, pending: %d)\n\n", total, doneCount, total-doneCount)

	stats := make([]*userStat, 0, len(userStats))
	for _, s := range userStats {
		stats = append(stats, s)
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].name < stats[j].name
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "User\tTotal\tDone\tPending")
	for _, stat := range stats {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n", stat.name, stat.total, stat.done, stat.pending)
	}
	if err := w.Flush(); err != nil {
		log.Fatalf("flush report: %v", err)
	}
}
