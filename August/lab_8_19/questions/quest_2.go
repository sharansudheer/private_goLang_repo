package questions

import ("fmt")

func quest_2(num1, num2 int) string{

        if num2%num1 == 0 {
                return fmt.Sprintf("%d is a Divisor of %d", num1,num2)
        } else {
                return fmt.Sprintf("%d is not a divisor of %d", num1, num2)
        }

}