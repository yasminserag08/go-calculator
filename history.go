package main

import "fmt"

var history []Calculation

func save(calc Calculation) {
	history = append(history, calc)
}

func viewHistory() {
	if len(history) == 0 {
		fmt.Println("No history yet")
		return
	}
	fmt.Println("************ History ************")
	for i, entry := range history {
		fmt.Printf("%d. %v %s %v = %v\n", i+1, entry.operand1, entry.operation, entry.operand2, entry.result)
	}
}

func clearHistory() {
	history = make([]Calculation, 0)
	fmt.Println("History cleared")
}
