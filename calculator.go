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

var operatorFuncs = map[string]func(float64, float64) (float64, error){
	"*": multiply,
	"/": divide,
	"+": add,
	"-": subtract,
}

func calculate(expression string) (float64, error) {
	tokens := tokenize(expression)
	postfix, err := getPostfix(tokens)
	if err != nil {
		return 0, err
	}
	result, err := computeResult(postfix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func tokenize(expression string) []string {
	expression = padExpression(expression)
	return strings.Fields(expression)
}

func getPostfix(tokens []string) ([]string, error) {
	var outputList []string
	var operatorStack []string
	for _, token := range tokens {
		if isOperator(token) {
			// if operator in stack has higher precedence over token, pop it to output list first
			for !isEmpty(operatorStack) && peek(operatorStack) != "(" && !precedes(token, peek(operatorStack)) {
				outputList = append(outputList, pop(&operatorStack))
			}
			push(&operatorStack, token)
		} else if isNumber(token) {
			outputList = append(outputList, token)
		} else if token == "(" {
			push(&operatorStack, token)
		} else if token == ")" {
			for !isEmpty(operatorStack) && peek(operatorStack) != "(" {
				outputList = append(outputList, pop(&operatorStack))
			}

			if isEmpty(operatorStack) {
				return nil, fmt.Errorf("mismatched parentheses: missing '('")
			}

			pop(&operatorStack)
		} else {
			return nil, fmt.Errorf("invalid token: %s", token)
		}

	}

	// add any remaining operators to the output list
	for !isEmpty(operatorStack) {
		if peek(operatorStack) == "(" {
			return nil, fmt.Errorf("mismatched parentheses: missing ')'")
		}

		outputList = append(outputList, pop(&operatorStack))
	}

	return outputList, nil
}

func computeResult(outputList []string) (float64, error) {
	var numberStack []float64
	for _, item := range outputList {
		if isNumber(item) {
			push(&numberStack, float(item))
		} else if isOperator(item) {
			if len(numberStack) < 2 {
				return 0, fmt.Errorf("malformed expression: missing numbers for operator %s", item)
			}
			operand2 := pop(&numberStack)
			operand1 := pop(&numberStack)
			operation := operatorFuncs[item]
			result, err := operation(operand1, operand2)
			if err != nil {
				return 0, err
			}
			push(&numberStack, result)
		}
	}

	if len(numberStack) != 1 {
		return 0, fmt.Errorf("malformed expression")
	}

	return pop(&numberStack), nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	_, ok := operatorPrecedence[token]
	return ok
}

func precedes(op1, op2 string) bool {
	return operatorPrecedence[op1] > operatorPrecedence[op2]
}

func float(number string) float64 {
	// ignored error because this func assumes that number has already been validated with isNumber
	num, _ := strconv.ParseFloat(number, 64)
	return num
}

func padExpression(expression string) string {
	var sb strings.Builder

	for i := 0; i < len(expression); i++ {
		char := expression[i]

		// if operator or parenthesis, pad with spaces
		switch char {
		case '+', '-', '*', '/', '(', ')':
			// add space if operator doesn't have space before it
			if sb.Len() > 0 && sb.String()[sb.Len()-1] != ' ' {
				sb.WriteByte(' ')
			}
			sb.WriteByte(char)
			// add space if operator doesn't have space after it
			if i+1 < len(expression) && expression[i+1] != ' ' {
				sb.WriteByte(' ')
			}
		default:
			sb.WriteByte(char)
		}
	}

	return sb.String()
}
