package questions
import "fmt"

func PrintName(rollNo int) {
	
    lastThreeDigits := rollNo % 1000
    if lastThreeDigits%2 == 0 {
        fmt.Println("Your last name is:")
    } else {
        fmt.Println("Your first name is:")
    }
}
