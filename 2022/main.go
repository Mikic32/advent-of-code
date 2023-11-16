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
			result += 1
		case "C":
			result += 2
		}
		result += 0
	case "Y":
		switch player1 {
		case "A":
			result += 1
		case "B":
			result += 2
		case "C":
			result += 3
		}
		result += 3
	case "Z":
		switch player1 {
		case "A":
			result += 2
		case "B":
			result += 3
		case "C":
			result += 1
		}
		result += 6
	}
	return result, true
}
