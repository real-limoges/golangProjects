package main

import (
    "fmt"
)

func main() {
    coins := []int{200, 100, 50, 20, 10, 5, 2, 1}
    //val := 20
    for i:= 0; i < len(coins); i++ {
        fmt.Println(coins[i])
    }
    fmt.Println(addChange(coins))
}

func addChange(coins []int) (bool) {
    sum := 0
    for i := 0; i < len(coins); i++ {
        sum += coins[i]
    }
    if sum == 200 {
        return true
    }
    return false
}
