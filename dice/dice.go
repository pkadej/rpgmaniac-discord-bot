package dice

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var diceRegexp = regexp.MustCompile(`(\d+)?k(2|4|6|8|12|20)([\+\-]\d+)?`)

const diceSeparator = "k"

func IsDiceMessage(message string) bool {
	return diceRegexp.MatchString(message)
}
func CalculateDices(rollString string) int {
	extractedK := strings.Split(rollString, diceSeparator)
	multiplier := 1
	addExtra := 0

	if val, err := strconv.Atoi(extractedK[0]); err == nil {
		multiplier = val
	}

	var diceRange int

	if strings.Contains(extractedK[1], "+") {
		plussplit := strings.Split(extractedK[1], "+")
		if val, err := strconv.Atoi(plussplit[1]); err == nil {
			addExtra = val
		}

		if val, err := strconv.Atoi(plussplit[0]); err == nil {
			diceRange = val
		}
	} else if strings.Contains(extractedK[1], "-") {
		minussplit := strings.Split(extractedK[1], "-")
		if val, err := strconv.Atoi(minussplit[1]); err == nil {
			addExtra = (-1) * val
		}

		if val, err := strconv.Atoi(minussplit[0]); err == nil {
			diceRange = val
		}
	} else {
		if val, err := strconv.Atoi(extractedK[1]); err == nil {
			diceRange = val
		}
	}

	return multiplier*(rand.Intn(diceRange)+1) + addExtra
}
