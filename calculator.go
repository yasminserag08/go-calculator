package main

import (
	"fmt"
	"strconv"
	"strings"
)

var operatorPrecedence = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
}

var operatorFuncs = map[string]func(float64, float64) float64{
	"*": multiply,
	"/": divide,
	"+": add,
	"-": subtract,
}

func parse(expression string) {
	tokens := tokenize(expression)
	fmt.Println(tokens)
	postfix := getPostfix(tokens)
	fmt.Println(postfix)
}

func tokenize(expression string) []string {
	return strings.Fields(expression)
}

func getPostfix(tokens []string) []string {
	var outputList []string
	var operatorStack []string
	for _, token := range tokens {
		if isOperator(token) {
			// if operator in stack has higher precedence over token, pop it to output list first
			for !isEmpty(operatorStack) && !precedes(token, peek(operatorStack)) {
				fmt.Printf("%s precedes %s\n", token, peek(operatorStack))
				outputList = append(outputList, pop(&operatorStack))
			}
			push(&operatorStack, token)
		} else if isNumber(token) {
			outputList = append(outputList, token)
		}
	}

	// add any remaining operators to the output list
	for !isEmpty(operatorStack) {
		outputList = append(outputList, pop(&operatorStack))
	}

	return outputList
}

func divide(a, b float64) float64 {
	return a / b
}

func multiply(a, b float64) float64 {
	return a * b
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	_, ok := operatorPrecedence[token]
	return ok
}

func push(stack *[]string, item string) {
	*stack = append(*stack, item)
}

func peek(stack []string) string {
	return stack[len(stack)-1]
}

func pop(stack *[]string) string {
	last := len(*stack) - 1
	item := (*stack)[last]
	*stack = (*stack)[:last]
	fmt.Printf("Popped %s\n", item)
	return item
}

func precedes(op1, op2 string) bool {
	return operatorPrecedence[op1] > operatorPrecedence[op2]
}

func isEmpty(slice []string) bool {
	return len(slice) == 0
}
