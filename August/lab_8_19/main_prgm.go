//1. Write a program in Go to reverse a string
//CB.SC.P2CSE24010

package main

import (
	"fmt"
	"lab_8_19/questions"
	"sync"
)


func main() {

	var check bool = true
	var quest int
	fmt.Println("Input")
	fmt.Scanln(&quest)
	for check {

		if quest == 0 {

			fmt.Println("Nice Day")
			check = false

		} else if quest> 7 {

			fmt.Println("Wrong Input")

			fmt.Println("Input")
			fmt.Scanln(&quest)
		
		} else {
			switch quest {

				case 1:{
					var sentence string
					fmt.Println("Enter Your string")
					fmt.Scanln(&sentence)

		
					reverse_str := questions.ReverseString(sentence)

					fmt.Println(reverse_str)

					fmt.Println("Input")
					fmt.Scanln(&quest)
				}

				case 2:{
					var num1, num2 int
					
					fmt.Print("Enter the first number: ")
					fmt.Scan(&num1)
					fmt.Print("Enter the second number: ")
					fmt.Scan(&num2)

					questions.Quest_2(num1,num2)

					fmt.Println("Input")
					fmt.Scanln(&quest)
				}

				case 3:{
					var rollNo int
					fmt.Print("Enter your roll number: ")
					fmt.Scan(&rollNo)
				
					questions.PrintName(rollNo)
				
				}
				
				/*
				Question 4
				 Returns the number of odd and even numbers in the array
				*/
				case 4:{
					var n int
					fmt.Print("Enter the size of the array: ")
					fmt.Scan(&n)

					arr := make([]int, n)
					fmt.Println("Enter the elements of the array:")
					for i := 0; i < n; i++ {
						fmt.Scan(&arr[i])
					}

					odd, even := questions.Find_odd_even(arr)
					fmt.Printf("Number of odd numbers: %d\n", odd)
					fmt.Printf("Number of even numbers: %d\n", even)

					}
				case 5:{
					var n int
						fmt.Print("Enter the size of the array: ")
						fmt.Scan(&n)

						arr := make([]int, n)
						fmt.Println("Enter the elements of the array:")
						for i := 0; i < n; i++ {
							fmt.Scan(&arr[i])
						}

						var wg sync.WaitGroup
						wg.Add(3)

						go func() {
							questions.ReverseArray(arr)
							fmt.Println("Reversed array:", arr)
							wg.Done()
						}()

						go func() {
							odd, even := questions.CountOddEven(arr)
							fmt.Printf("Number of odd numbers: %d\n", odd)
							fmt.Printf("Number of even numbers: %d\n", even)
							wg.Done()
						}()

						go func() {
							zeros := questions.CountZeros(arr)
							fmt.Printf("Number of zeros: %d\n", zeros)
							wg.Done()
						}()

						wg.Wait()
				}
					case 6:{

						var wg sync.WaitGroup
				wg.Add(2)
				
				{

				go func() {
					name := questions.Get_Name()
					wg.Done()
					questions.Print_data(name, questions.Get_rollno())
				}()

				go func() {
					rollno := questions.Get_rollno()
					wg.Done()
					questions.Print_data(questions.Get_Name(), rollno)
				}()

				wg.Wait()
				}
						
					}
				case 7:{
					var r questions.Rect
					r.X = 5
					r.Y = 10
				
					questions.CalculateArea(&r)
					questions.CalculatePerimeter(&r)
				
					fmt.Printf("Area: %.2f\n", r.Area)
					fmt.Printf("Perimeter: %.2f\n", r.Perimeter)
					
			}

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
you'll need to expose them through public functions or use reflection. */