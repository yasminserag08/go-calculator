package main

import "fmt"

type HistoryEntry struct {
	expression string
	result     float64
}

var history []HistoryEntry

func save(entry HistoryEntry) {
	history = append(history, entry)
}

func viewHistory() {
	if len(history) == 0 {
		fmt.Println("No history yet")
		return
	}
	fmt.Println("************ History ************")
	for i, entry := range history {
		fmt.Printf("%d. %s = %v\n", i+1, entry.expression, entry.result)
	}
}

func clearHistory() {
	history = make([]HistoryEntry, 0)
	fmt.Println("History cleared")
}
