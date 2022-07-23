package main

import "fmt"

const text string = "Text ini adalah constant"
const status bool = false


func main() {
    const info = "tidak dapat dirubah"
    fmt.Println(text, status, info)
}
