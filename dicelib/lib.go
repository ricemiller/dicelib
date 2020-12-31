/*
	Requirements:
1d10 = rolls 1 10-sided die
2d10 = rolls 2 10-sided dice
1d10+1 = rolls 1 10-sided die and adds 1
1d10+1d4 = rolls 1 10-sided die and adds 1 4-sided die


*/

package dice

import (
	"math/rand"
	"time"
	"strings"
	"strconv"
)

// MakeRoll returns the result of a die roll
func MakeRoll(dieType int) int {
	rand.Seed(time.Now().UnixNano())
	return (int(rand.Uint64()) % dieType) + 1
}

// DummyRoll always returns average (for testing purposes)
func DummyRoll(dieType int) int {
	return dieType/2
}


// SimpleRoll makes a simple roll
func SimpleRoll(roll string) int {
	var result int

	components := strings.Split(roll, "d")

	if len(components) > 1 {

		numDice, _ := strconv.Atoi(components[0])

		dieType, _ := strconv.Atoi(components[1])

		for i := 0; i < numDice; i++ {
			result += MakeRoll(dieType)
			//result += DummyRoll(dieType)
		}
	} else {
		result, _ = strconv.Atoi(components[0])
	}

	return result
}

// Split consumes the initial command and splits it into positive and negative values
func Split(command string) ([]string, []string) {
	roll := ""
	var positives []string
	var negatives []string

	for i := len(command)-1 ; i >= 0 ; i-- {
		if command[i] == '+' {
			positives = append(positives, roll)
			roll = ""
		} else if command[i] == '-' {
			negatives = append(negatives, roll)
			roll = ""
		} else {
			roll = string(command[i]) + roll
		}
	}
	positives = append(positives, roll)
	return positives, negatives
}

// Roll splits the command into smaller chunks for SimpleRoll
func Roll(commands string) int {
	var result int
	var positive, negative []string

	positive, negative = Split(commands)

	for _, roll := range positive {
		result += SimpleRoll(roll)
	}

	for _, roll := range negative {
		result -= SimpleRoll(roll)
	}

	return result
}
