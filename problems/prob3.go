package main

import "fmt"

func main() {
    numbers := []int{1, 3, 5, 7, 9, 2}
    fmt.Println(secondLargest(numbers))
}

func secondLargest(numbers []int) int {
    max1 := numbers[0]
    max2 := -1
    for i := 1; i < len(numbers); i++ {
        if numbers[i] > max1 {
            max2 = max1
            max1 = numbers[i]
        } else if numbers[i] > max2 {
            max2 = numbers[i]
        }
    }
    return max2
}
