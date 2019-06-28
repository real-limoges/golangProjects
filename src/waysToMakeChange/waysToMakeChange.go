package main

import (
    "fmt"
)

func main() {
    const total int = 200
    coins := [...]int{1, 2, 5, 10, 20, 50, 100, 200}

    var ways [total+1]int
    ways[0] = 1

    for i := 0; i < len(coins); i++ {
        for j := 0; j < len(ways) - coins[i]; j++ {
            ways[j + coins[i]] += ways[j]
        }
    }
    fmt.Println(ways[total])
}

