package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to my simple calculator")
	reader := bufio.NewReader(os.Stdin)
	// infinite loop until user chooses to exit
outer:
	for {
		displayMenu()

		fmt.Print("Option: ")
		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error occurred while reading the option: ", err)
			continue
		}

		option = strings.TrimSpace(option)

		err = validateOption(option)

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch option {
		case "1":
			fmt.Print("Enter expression: ")
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("an error occurred while reading the expression: ", err)
			} else {
				result, err := calculate(input)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Result:", result)
					save(HistoryEntry{input, result})
				}
			}

		case "2":
			viewHistory()
			continue
		case "3":
			clearHistory()
		case "4":
			fmt.Println("Exiting...")
			break outer
		}
	}
}
