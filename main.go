package main

import "fmt"

func main() {
	fmt.Println("Welcome to my simple calculator")

	// infinite loop until user chooses to exit
outer:
	for {
		var number1, number2 float64
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&number1) // scanln is enough for now since no whitespaces are needed in the input
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&number2)

		var operation string
		fmt.Print("Enter operation (+, -, *, /): ")
		fmt.Scanln(&operation)

		switch operation {
		case "+":
			fmt.Println("Result: ", number1+number2)
		case "-":
			fmt.Println("Result: ", number1-number2)
		case "*":
			fmt.Println("Result: ", number1*number2)
		case "/":
			if number2 == 0 {
				fmt.Println("Error: Division by zero")
				continue
			}
			fmt.Println("Result: ", number1/number2)
		default:
			fmt.Println("Invalid operation")
			continue
		}

		// after printing the result, ask the user if they want to exit
		var choice string
		fmt.Println("Do you want to exit? y / n")
		fmt.Scanln(&choice)

		switch choice {
		case "y":
			break outer
		case "n":
			continue
		default:
			fmt.Println("Invalid choice, continuing.") // if neither y or n, default is to continue
		}
	}
}
