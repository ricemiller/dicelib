Dicelib is a Go library for rolling dice, tabletop RPG style.

The **Roll()** function takes a string and returns an Integer.

## Syntax
- 1d10 -> Rolls one 10-sided die
- 1d10+4 -> Rolls one 10-sided die and adds 4 to the result
- 1d10-1d4 -> Rolls one 10-sided die and one 4-sided die and subtracts their results together
- 4d6h3 -> Rolls four 6-sided dice, keeps the 3 highest results and adds them together 
- 4d6l3 -> Rolls four 6-sided dice, keeps the 3 lowest results and adds them together 

All these operators can be combined to create complex rolls:
`-4d6h3+1d10-8+2d10l1`

## How to use
Here's a quick example that will take an argument, roll dice and give the output.

``` go
package main

import (
	"dicelib/dice"
	"fmt"
	"os"
)

func main() {
	dice.Init() //Seed the random library
	command := os.Args[1] //Read command-line argument
	fmt.Println(dice.Roll(command)) // Roll and print result
}
```

```
~/go/src/dicelib go run main.go 1d10
8
```

