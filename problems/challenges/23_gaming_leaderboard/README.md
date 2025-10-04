# Challenge #23: Gaming Leaderboard

## Objective
Write a Go program that takes a map representing player scores and generates a sorted leaderboard. This is a common requirement in gaming or any system that ranks users.

## Difficulty
Medium

## Concepts Tested
* Maps (`map[string]int`)
* Structs (`struct`)
* Slices and Sorting (`sort.Slice`)
* Functions

## Rules/Logic
1.  Start with a map where keys are player names (string) and values are their scores (int).
2.  Define a struct named `Player` with two fields: `Name` (string) and `Score` (int).
3.  Create a function `createLeaderboard(scores map[string]int) []Player`.
4.  Inside this function, you must convert the input map into a slice of `Player` structs.
5.  After converting, sort the slice of `Player` structs in **descending order** of score. If two players have the same score, their relative order doesn't matter.
6.  The function should return the final, sorted slice of `Player`s.
7.  In your `main` function:
    * Create the initial map of player scores.
    * Call your `createLeaderboard` function.
    * Loop through the returned leaderboard slice and print each player's rank, name, and score.

## Your Tasks
1.  Create a new directory: `problems/challenges/23_gaming_leaderboard/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
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

// Expected Terminal Output (Alice and Charlie can be in any order):
--- Official Leaderboard ---
1. alice - 100
2. charlie - 100
3. shivam - 95
4. bob - 85