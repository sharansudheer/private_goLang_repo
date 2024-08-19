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

   for  check == true{
	switch check == true {
		
		
	}
	fmt.Println("Enter Your ") 
  

   }
}