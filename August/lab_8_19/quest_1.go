//1. Write a program in Go to reverse a string
//CB.SC.P2CSE24010

package main

import "fmt"
func reverseString(str string) string{
   byte_str := []rune(str)
   for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
      byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
   }
   return string(byte_str)
}

func main(){

	var check bool = true
	var questions int

	for check == true{
		
		if questions == 0{
			break
		}else if questions > 7{

			fmt.Println("Wrong Input")
			
		}else{

			switch questions{

				case 1: 
					var sentence string
					fmt.Println("Enter Your string") 
					fmt.Scanln(&sentence)
					fmt.Println(reverseString(sentence))
	
				

			case 2:



				

			}

		}	
   }
}