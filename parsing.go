package main

import "strings"

func tokenize(expression string) []string {
	expression = padExpression(expression)
	return strings.Fields(expression)
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
