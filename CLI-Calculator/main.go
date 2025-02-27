package main

import "fmt"

func main() {

	var num1, num2 int
	var operator string

	fmt.Print("Enter a number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter another number: ")
	fmt.Scan(&num2)

	fmt.Println("Enter operator (+, -, *, /)")

	fmt.Scan(&operator)
	result := performOperation(num1, num2, operator)
	fmt.Println(result)

}
func performOperation(num1 int, num2 int, operator string) string {

	switch operator {
	case "+":
		return fmt.Sprintf("Result: %d", num1+num2)
	case "-":
		return fmt.Sprintf("Result: %d", num1-num2)
	case "*":
		return fmt.Sprintf("Result: %d", num1*num2)
	case "/":
		if num2 == 0 {
			return fmt.Sprintf("Cannot divide by zero")
		}
		return fmt.Sprintf("Result: %d", num1/num2)
	default:
		return "Invalid operator"
	}
}
