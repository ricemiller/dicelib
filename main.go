package main

import (
	"dicelib/dice"
	"fmt"
	"os"
)

func main() {
	dice.Init()
	command := os.Args[1]
	fmt.Println(dice.Roll(command))
}
