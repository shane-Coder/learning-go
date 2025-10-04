package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name  string
	Score int
}

func createLeaderboard(scores map[string]int) []Player {
	var leaderboard []Player
	for name, score := range scores {
		leaderboard = append(leaderboard, Player{Name: name, Score: score})
	}
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})
	return leaderboard
}

func main() {
	playerScores := map[string]int{
		"shivam":  95,
		"alice":   100,
		"bob":     85,
		"charlie": 100,
	}

	leaderboard := createLeaderboard(playerScores)

	fmt.Println("--- Official Leaderboard ---")
	for rank, player := range leaderboard {
		fmt.Printf("%d. %s - %d\n", rank+1, player.Name, player.Score)
	}
}
