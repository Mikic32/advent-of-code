package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("strategyGuide.txt")
	if err != nil {
		fmt.Println("error loading file:", err)
		return
	}
	contentStr := string(content)
	splitedContent := strings.Split(contentStr, "\n")
	score := 0
	for _, round := range splitedContent {
		if v, ok := roundResult(round); ok {
			score += v
		}
	}
	println(score)
}

func roundResult(round string) (int, bool) {
	players := strings.Split(round, " ")
	if len(players) != 2 {
		return 0, false
	}
	result := 0
	player1, player2 := players[0], players[1]
	switch player2 {
	case "X":
		switch player1 {
		case "A":
			result += 3
		case "B":
			result += 0
		case "C":
			result += 6
		}
		result += 1
	case "Y":
		switch player1 {
		case "A":
			result += 6
		case "B":
			result += 3
		case "C":
			result += 0
		}
		result += 2
	case "Z":
		switch player1 {
		case "A":
			result += 0
		case "B":
			result += 6
		case "C":
			result += 3
		}
		result += 3
	}
	return result, true
}

