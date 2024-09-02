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
						sum := findSum(arr)
						fmt.Println("Sum of the array:", sum)
						wg.Done()
					}()

					go func() {
						average := findAverage(arr)
						fmt.Println("Average of the array:", average)
						wg.Done()
					}()

					go func() {
						max := findMax(arr)
						fmt.Println("Maximum element in the array:", max)
						wg.Done()
					}()

					wg.Wait()
				}
					case 6:{
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
							reverseArray(arr)
							fmt.Println("Reversed array:", arr)
							wg.Done()
						}()

						go func() {
							odd, even := countOddEven(arr)
							fmt.Printf("Number of odd numbers: %d\n", odd)
							fmt.Printf("Number of even numbers: %d\n", even)
							wg.Done()
						}()

						go func() {
							zeros := countZeros(arr)
							fmt.Printf("Number of zeros: %d\n", zeros)
							wg.Done()
						}()

						wg.Wait()
					}
				case 7:{
					var wg sync.WaitGroup
				wg.Add(2)
				
				{

				go func() {
					name := Get_Name()
					wg.Done()
					print_data(name, Get_rollno())
				}()

				go func() {
					rollno := Get_rollno()
					wg.Done()
					print_data(Get_Name(), rollno)
				}()

				wg.Wait()
				}
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