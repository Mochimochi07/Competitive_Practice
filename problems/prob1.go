package main

import "fmt"

func main() {
    numbers := []int{2, 4, 6, 2, 5}
    fmt.Println(largestSum(numbers))
}

func largestSum(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    } else if len(numbers) == 1 {
        return numbers[0]
    }

    first := numbers[0]
    second := max(numbers[0], numbers[1])
    for i := 2; i < len(numbers); i++ {
        current := max(first + numbers[i], second)
        first = second
        second = current
    }

    return second
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
