package main

import "fmt"

func displayMenu() {
	fmt.Println("*********************************")
	fmt.Println("Select one of the following options:")
	fmt.Println("1: Calculate")
	fmt.Println("2: View History")
	fmt.Println("3: Clear History")
	fmt.Println("4: Exit")
	fmt.Println("*********************************")
}

func validateOption(option int) error {
	if option < 1 || option > 4 {
		return fmt.Errorf("invalid option")
	}
	return nil
}
