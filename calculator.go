package main

import "fmt"

func calculate(calc *Calculation) error {
	switch calc.operation {
	case "+":
		calc.result = calc.operand1 + calc.operand2
	case "-":
		calc.result = calc.operand1 - calc.operand2
	case "*":
		calc.result = calc.operand1 * calc.operand2
	case "/":
		if calc.operand2 == 0 {
			return fmt.Errorf("cannot divide by zero")
		}
		calc.result = calc.operand1 / calc.operand2
	default:
		return fmt.Errorf("invalid operation")
	}
	return nil
}
