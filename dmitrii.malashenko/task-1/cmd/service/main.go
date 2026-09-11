package main

import "fmt"

func main() {
	var a, b int
	var op string

	_, errA := fmt.Scan(&a)

	if errA != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, errB := fmt.Scan(&b)

	if errB != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, errOp := fmt.Scan(&op)

	if errOp != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
		return
	}
}
