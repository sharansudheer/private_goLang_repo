package main

import "fmt"

func quest_2() {
        var num1, num2 int

        fmt.Print("Enter the first number: ")
        fmt.Scan(&num1)

        fmt.Print("Enter the second number: ")
        fmt.Scan(&num2)

        if num2%num1 == 0 {
                fmt.Println(num1, " is a divisor of ", num2)
        } else {
                fmt.Println(num1, " is not a divisor of ", num2)
        }
}