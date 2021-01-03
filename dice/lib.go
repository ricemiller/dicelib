package dice

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"sort"
)

// MakeRoll returns the result of a die roll
func MakeRoll(dieType int) int {
	return (rand.Intn(dieType) + 1)
}

// DummyRoll always returns average (for testing purposes)
func DummyRoll(dieType int) int {
	return dieType / 2
}

// SimpleRoll makes a simple roll
func SimpleRoll(roll string) int {
	var result, dieType int
	var results []int

	if strings.Contains(roll, "d") {
		components := strings.Split(roll, "d")
		numDice, _ := strconv.Atoi(components[0])
		keepHigh := 0
		keepLow := 0

		if strings.Contains(components[1], "h") {
			comps := strings.Split(components[1], "h")
			dieType, _ = strconv.Atoi(comps[0])
			keepHigh, _ = strconv.Atoi(comps[1])
		} else if strings.Contains(components[1], "l") {
			comps := strings.Split(components[1], "l")
			dieType, _ = strconv.Atoi(comps[0])
			keepLow, _ = strconv.Atoi(comps[1])
		} else {
			dieType, _ = strconv.Atoi(components[1])
		}

		for i := 0; i < numDice; i++ {
			results = append(results, MakeRoll(dieType))
			//results = append(results, DummyRoll(dieType))
		}
		sort.Ints(results)
		if keepLow > 0 {
			for i := 0; i < keepLow; i++ {
				result += results[i]
			}
		} else if keepHigh > 0 {
			for i := 0; i < keepHigh; i++ {
				result += results[numDice-1-i]
			}

		} else {
			for _, res := range results {
				result += res
			}
		}

	} else {
		result, _ = strconv.Atoi(roll)
	}

	return result
}

// Split consumes the initial command and splits it into positive and negative values
func Split(command string) ([]string, []string) {
	roll := ""
	var positives []string
	var negatives []string

	for i := len(command) - 1; i >= 0; i-- {
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

// Init just sets the seed of the random
func Init() {
	rand.Seed(time.Now().UnixNano())
}
