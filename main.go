package main

import "fmt"

func main() {
	fmt.Println("Welcome to my simple calculator")

	// infinite loop until user chooses to exit
outer:
	for {
		displayMenu()

		var option int
		fmt.Print("Choice: ")
		fmt.Scanln(&option)

		err := validateOption(option)

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch option {
		case 1:
			var number1, number2 float64
			fmt.Print("Enter the first number: ")
			fmt.Scanln(&number1) // scanln is enough for now since no whitespaces are needed in the input
			fmt.Print("Enter the second number: ")
			fmt.Scanln(&number2)

			var operation string
			fmt.Print("Enter operation (+, -, *, /): ")
			fmt.Scanln(&operation)

			result, err := calculate(number1, number2, operation)

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(result)
		case 2:
			// TODO
			fmt.Println("View history - coming soon!")
			continue
		case 3:
			// TODO
			fmt.Println("Clear history - coming soon!")
		case 4:
			fmt.Println("Exiting...")
			break outer
		}
	}
}
