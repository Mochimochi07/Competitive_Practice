package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(reverseList(nums))
}

func reverseList(nums []int) []int {
    for i := 0; i < len(nums)/2; i++ {
        j := len(nums) - i - 1
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}
