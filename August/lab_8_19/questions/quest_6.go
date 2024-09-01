package questions

import (
    "fmt"
)

func Get_Name() string {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    return name
}

func Get_rollno() int {
    var rollno int
    fmt.Print("Enter your roll number: ")
    fmt.Scan(&rollno)
    return rollno
}

func Print_data(name string, rollno int) {
    fmt.Printf("Name: %s\nRoll No: %d\n", name, rollno)
}