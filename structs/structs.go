package structs

import (
	"strings"
)

type Game string

const (
	GameAlien       Game = "Alien RPG"
	GameTales       Game = "Tales from the loop"
	GameUnsupported Game = ""
)

func DetermineGame(game string) Game {
	if strings.Contains(game, "alien") {
		return GameAlien
	}

	if strings.Contains(game, "tales") {
		return GameTales
	}

	return GameUnsupported
}
