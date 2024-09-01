package questions



func Find_odd_even(arr []int) (int, int) {
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