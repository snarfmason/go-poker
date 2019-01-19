package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "go-poker/poker"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  line, err := reader.ReadString('\n')
  round := 0
  for err == nil {
    strs := strings.Split(line, "|")
    n_hands := len(strs)
    hands := make([]poker.Hand, n_hands)
    for i := 0; i < n_hands; i++ {
      hands[i] = poker.ParseHand(strs[i])
    }

    round++
    fmt.Println("\nRound", round)
    for i := 0; i < n_hands; i++ {
      fmt.Println(hands[i])
    }
    line, err = reader.ReadString('\n')
  }
}
