package main

import "fmt"

func main() {

    messages := make(chan string)
    messages1 := make(chan int)

    go func() { messages <- "ping" }()
    go func() { messages1 <- 1 }() 
    msg := <-messages
    fmt.Println(msg)
    msg1 := <-messages1
    fmt.Println(msg1)
}
