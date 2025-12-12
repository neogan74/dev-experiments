package main

import (
	"sort"
	"strconv"
	"strings"
)

// countMentions processes events in timestamp order, applying status changes
// before messages at the same timestamp.
func countMentions(numberOfUsers int, events [][]string) []int {
	type event struct {
		typ     string
		ts      int
		payload string
	}

	typePriority := func(t string) int {
		if t == "OFFLINE" {
			return 0
		}
		return 1
	}

	parsed := make([]event, len(events))
	for i, e := range events {
		ts, _ := strconv.Atoi(e[1])
		parsed[i] = event{
			typ:     e[0],
			ts:      ts,
			payload: e[2],
		}
	}

	sort.SliceStable(parsed, func(i, j int) bool {
		if parsed[i].ts != parsed[j].ts {
			return parsed[i].ts < parsed[j].ts
		}
		return typePriority(parsed[i].typ) < typePriority(parsed[j].typ)
	})

	mentions := make([]int, numberOfUsers)
	online := make([]bool, numberOfUsers)
	offlineUntil := make([]int, numberOfUsers) // timestamp when user comes back online
	for i := range online {
		online[i] = true
	}

	refresh := func(now int) {
		for i := 0; i < numberOfUsers; i++ {
			if !online[i] && now >= offlineUntil[i] {
				online[i] = true
			}
		}
	}

	prevTS := -1
	for _, ev := range parsed {
		if ev.ts != prevTS {
			refresh(ev.ts)
			prevTS = ev.ts
		}

		if ev.typ == "OFFLINE" {
			id, _ := strconv.Atoi(ev.payload)
			online[id] = false
			offlineUntil[id] = ev.ts + 60
			continue
		}

		for _, tok := range strings.Fields(ev.payload) {
			switch tok {
			case "ALL":
				for i := range mentions {
					mentions[i]++
				}
			case "HERE":
				for i := 0; i < numberOfUsers; i++ {
					if online[i] {
						mentions[i]++
					}
				}
			default:
				if strings.HasPrefix(tok, "id") {
					if id, err := strconv.Atoi(tok[2:]); err == nil && id >= 0 && id < numberOfUsers {
						mentions[id]++
					}
				}
			}
		}
	}

	return mentions
}
