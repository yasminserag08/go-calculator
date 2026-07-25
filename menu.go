package main

import (
	"fmt"
	"strconv"
)

func displayMenu() {
	fmt.Println("*********************************")
	fmt.Println("Select one of the following options:")
	fmt.Println("1: Calculate")
	fmt.Println("2: View History")
	fmt.Println("3: Clear History")
	fmt.Println("4: Exit")
	fmt.Println("*********************************")
}

func validateOption(option string) error {

	optionInt, err := strconv.Atoi(option)
	if err != nil {
		return fmt.Errorf("invalid option: please enter a number")
	}

	if optionInt < 1 || optionInt > 4 {
		return fmt.Errorf("invalid option: choose between 1 and 4")
	}

	return nil
}
