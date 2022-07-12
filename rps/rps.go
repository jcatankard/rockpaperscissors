package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var (
	GAME         = Game{} // GAME is a global variable, it is initiated when webapp launched
	mapConstants = map[int]string{
		ROCK:     "rock",
		PAPER:    "paper",
		SCISSORS: "scissors",
	}
	mapStrings = map[string]int{
		"rock":     ROCK,
		"paper":    PAPER,
		"scissors": SCISSORS,
	}
)

// Round is a placeholder to return the results from the latest round
type Round struct {
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

// Game is a placeholder to keep track of the computer and player scores
type Game struct {
	ComputerScore int `json:"computer_score"`
	PlayerScore   int `json:"player_score"`
}

// updateScores updates the scores based on who won the round
func (g *Game) updateScores(roundResult string) {
	if roundResult == "player wins" {
		g.PlayerScore++
	} else if roundResult == "computer wins" {
		g.ComputerScore++
	}
}

// PlayRound orchestrates a new round
func PlayRound(playerChoice string) Round {

	playerValue := mapStrings[playerChoice]
	computerValue := computerPlays()
	roundResult := evaluateWinner(playerValue, computerValue)

	GAME.updateScores(roundResult)

	result := Round{
		ComputerChoice: mapConstants[computerValue],
		RoundResult:    roundResult,
	}
	return result
}

// computerPlays makes a randomly generated turn for the computer
func computerPlays() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

// evaluateWinner tells us if player wins, computer wins or draw
func evaluateWinner(p, c int) string {

	if p == c {
		return "draw"
	} else if p == (c+1)%3 {
		return "player wins"
	} else {
		return "computer wins"
	}
}
