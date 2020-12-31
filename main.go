package main

import (
	"dicelib/dicelib"
	"fmt"
	"os"
)

func main() {

	command := os.Args[1]
	fmt.Println(dice.Roll(command))

}
