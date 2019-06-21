package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  noSwitchWins := 0
  switchWins := 0
  total := 0

  for i := 0; i < 1000; i++ {
    car := r1.Intn(3)
    yourPick := r1.Intn(3)
    monty := r1.Intn(3)

    for {
      if monty != yourPick && monty != car {
        break
      }
      monty = r1.Intn(3)
    }

    if car == yourPick {
      noSwitchWins += 1
    } else {
      switchWins += 1
    }
    total += 1
  }

  noSwitchWinsPct := float64(noSwitchWins) / float64(total)
  switchWinsPct := float64(switchWins) / float64(total)

  fmt.Printf("If you didn't switch you'd win %.2f percent of the time\n",
             noSwitchWinsPct*100)
  fmt.Printf("If you did switch you'd win %.2f percent of the time\n",
             switchWinsPct*100)
}
