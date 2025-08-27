/*
Problem Statement #2: Simple Number Guessing Game

Objective:
	Write a Go program where you have a secretNumber, and you try to guess it. The program should loop, providing hints ("Too high" or "Too low") until the guess is correct.

Rules/Logic:

	1. Create two variables: secretNumber := 7 and guess := 0.

	2. Use a for loop that continues as long as the guess is not equal to the secretNumber.

	3. Inside the loop, you will simulate a new guess on each iteration. To keep it simple, you can just increment your guess variable by 1 in each loop run (guess++).

	4. Inside the loop, use an if/else statement to check if the current guess is lower or higher than the secretNumber and print the appropriate hint.

	5. Once the guess is correct, the loop will terminate. After the loop, print a success message.

--- Terminal Output ---

	// Test Case 1: Guessing down to the secret number
	Guess: 10 - Too high
	Guess: 9 - Too high
	Guess: 8 - Too high
	Guess: 7 - Correct! You've guessed the secret number.

	// Test Case 2: Guessing up from a lower number
	Guess: 10 - Too low
	Guess: 11 - Too low
	...
	Guess: 17 - Correct! You've guessed the secret number.

	// Test Case 3: Guessing up from a negative number
	Guess: -10 - Too low
	Guess: -9 - Too low
	...
	Guess: 1 - Correct! You've guessed the secret number.
*/

package main

import "fmt"

func main() {
	secretNumber := 10
	guess := 20

	for guess != secretNumber {
		temp := guess
		if temp < secretNumber {
			fmt.Printf("Guess: %d - Too low\n", temp)
			temp++
		} else if temp > secretNumber {
			fmt.Printf("Guess: %d - Too high\n", temp)
			temp--
		} else {
			break
		}
		guess = temp
	}
	fmt.Printf("Guess: %d - Correct! You've guessed the secret number.\n", guess)
}
