package main

import (
    "fmt"
    "math/rand"
)

func CalculateValue(values chan int) {
    value := rand.Intn(10)
    fmt.Println("Calculated Random Value: {}", value)
    values <- value
}

func main() {
    fmt.Println("Go Channel Tutorial")

  values := make(chan int,2)
  defer close(values)

    go CalculateValue(values)
    go CalculateValue(values)	
    value := <-values
    fmt.Println(value)
}
