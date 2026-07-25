package main

func push[T interface{}](stack *[]T, item T) {
	*stack = append(*stack, item)
}

func peek(stack []string) string {
	return stack[len(stack)-1]
}

func pop[T interface{}](stack *[]T) T {
	last := len(*stack) - 1
	item := (*stack)[last]
	*stack = (*stack)[:last]
	return item
}

func isEmpty(slice []string) bool {
	return len(slice) == 0
}
