package main

import "fmt"

func main() {
    words := []string{"cat", "dog", "cat", "bird", "cat"}
    target := "cat"
    fmt.Println(wordCount(words, target))
}

func wordCount(words []string, target string) int {
    count := 0
    for _, word := range words {
        if word == target {
            count++
        }
    }
    return count
}
