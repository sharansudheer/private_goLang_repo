package questions

func ReverseArray(arr []int) {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}

func CountOddEven(arr []int) (int, int) {
    oddCount, evenCount := 0, 0
    for _, num := range arr {
        if num%2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }
    return oddCount, evenCount
}

func CountZeros(arr []int) int {
    zeroCount := 0
    for _, num := range arr {
        if num == 0 {
            zeroCount++
        }
    }
    return zeroCount
}