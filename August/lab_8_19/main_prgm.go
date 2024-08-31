//1. Write a program in Go to reverse a string
//CB.SC.P2CSE24010

package main

import (
	"fmt"
	"lab_8_19/questions"
)

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func main() {

	var check bool = true
	var quest int
	fmt.Println("Input")
	fmt.Scanln(&quest)
	for check == true {

		if quest == 0 {
			fmt.Println("Nice Day")
			break
		} else if quest> 7 {

			fmt.Println("Wrong Input")
			break
		} else {

			switch quest {

			case 1:
				var sentence string
				fmt.Println("Enter Your string")
				fmt.Scanln(&sentence)
				fmt.Println(reverseString(sentence))

			case 2:
				var num1, num2 int
				
				fmt.Print("Enter the first number: ")
				fmt.Scan(&num1)

				fmt.Print("Enter the second number: ")
				fmt.Scan(&num2)
				var question string
				question = questions.Quest_2(8,2)
				fmt.Print(question)
				
			}

		}
	}
}
//do not use :=, it causes type declaration issues. quest_2 Is 
//outputing String
/* If the external package is in a different directory, 
you might need to specify the path when importing it.
For larger projects, 
consider using Go modules to manage dependencies and package organization.
If you want to access private functions (starting with a lowercase letter)
from the external package, 
you'll need to expose them through public functions or use reflection.*/