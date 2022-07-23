package main

import "fmt"

func TesterReturn(text string) string {
    return text
}

func TesterReturn2(num int) int {
    return num
}

func main() {
    text := TesterReturn("Return Text and number")
    num := TesterReturn2(24)
    fmt.Println(text, num)
}
