package main

import ("thisisatest/add"
		"fmt"
		"thisisatest/div"
		"thisisatest/multi"
		"thisisatest/sub"
)

func main(){
	var a, b float64
	var oper string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&oper)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	switch oper{
	case "+":
		add.PrintDetails(a,b)
	case "/":
		div.PrintDetails(a,b)
	case "*":
		multi.PrintDetails(a,b)
	case "-":
		sub.PrintDetails(a,b)
		


	}

}