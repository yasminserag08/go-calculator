package main

import (
	"bufio"
	"fmt"
	"os"
)

type Calculation struct {
	operand1  float64
	operand2  float64
	operation string
	result    float64
}

func main() {
	fmt.Println("Welcome to my simple calculator")

	// infinite loop until user chooses to exit
outer:
	for {
		displayMenu()

		var option int
		fmt.Print("Option: ")
		fmt.Scanln(&option)

		err := validateOption(option)

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch option {
		case 1:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter expression: ")
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("an error occurred while reading the expression: ", err)
			} else {
				parse(input)
			}

		case 2:
			viewHistory()
			continue
		case 3:
			clearHistory()
		case 4:
			fmt.Println("Exiting...")
			break outer
		}
	}
}
