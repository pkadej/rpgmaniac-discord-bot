package dice

import (
	"fmt"
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

func parseRollString(rollString string) (multiplier int, diceRange int, addExtra int) {

	extractedK := strings.Split(rollString, diceSeparator)
	multiplier = 1
	addExtra = 0

	if val, err := strconv.Atoi(extractedK[0]); err == nil {
		multiplier = val
	}

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

	return
}

func CalculateDices(rollString string) int {
	multiplier, diceRange, addExtra := parseRollString(rollString)
	diceResults := 0
	for iteration := 0; iteration < multiplier; iteration++ {
		diceResults += rand.Intn(diceRange) + 1
	}

	return diceResults + addExtra
}

func DescribeDices(message string) string {
	multiplier, diceRange, addExtra := parseRollString(message)
	result := 0
	successes := 0
	var sb strings.Builder

	for iteration := 0; iteration < multiplier; iteration++ {
		randomResult := rand.Intn(diceRange) + 1
		result += randomResult
		if randomResult == 6 {
			successes++
		}
		rollText := fmt.Sprintf("%s%d[**%d**] ", diceSeparator, diceRange, randomResult)
		sb.WriteString(rollText)

	}
	if addExtra != 0 {
		if addExtra > 0 {
			sb.WriteString("+")
		}

		sb.WriteString(fmt.Sprintf("%d", addExtra))
		result += addExtra
	}

	sb.WriteString(fmt.Sprintf(" Total=[%d]", result))
	if successes > 0 {
		sb.WriteString(fmt.Sprintf("  **SUCCESSES=[%d]**", successes))
	}

	return sb.String()
}
