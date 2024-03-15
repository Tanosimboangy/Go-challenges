func PositiveSum(numbers []int) (sum int) {
    for i := 0; i < len(numbers); i++ {
        if (numbers[i] > 0) {sum += numbers[i]}
    }
    return
}

// I chose this solution because it caught my eyes when I see it return like just nothing