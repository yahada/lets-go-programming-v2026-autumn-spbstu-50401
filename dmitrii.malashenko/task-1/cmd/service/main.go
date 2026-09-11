package main

import "fmt"

func main() {
	var a, b int
	var op string

	_, err_a := fmt.Scan(&a)

	if err_a != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err_b := fmt.Scan(&b)

	if err_b != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err_op := fmt.Scan(&op)

	if err_op != nil {
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
