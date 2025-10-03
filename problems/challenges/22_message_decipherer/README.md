# Challenge #22: Message Decipherer

## Objective
Write a Go program to decipher a message that has been encoded by reversing each word in a sentence. This challenge will test your ability to correctly handle Unicode characters.

## Difficulty
Medium

## Concepts Tested
* Strings and Runes (`string`, `[]rune`)
* String Manipulation (`strings.Fields`, `strings.Builder`)
* Slices and Loops
* Functions

## Rules/Logic
1.  Create a helper function `reverseWord(word string) string`. This function must correctly reverse a string that might contain multi-byte Unicode characters (like `é` or `🚀`). To do this properly, you must:
    * Convert the string to a slice of runes (`[]rune`).
    * Reverse the rune slice.
    * Convert the reversed slice back to a string.
2.  Create a main function `decipherMessage(message string) string`.
3.  Inside `decipherMessage`, first split the input `message` into a slice of words using `strings.Fields()`.
4.  Loop through the slice of words. For each word, call your `reverseWord` helper function to get the decoded word.
5.  Use a `strings.Builder` to efficiently join the decoded words back into a single sentence, separated by spaces.
6.  In your `main` function, test your `decipherMessage` function with a few encoded sentences, including one with multi-byte characters.

## Your Tasks
1.  Create a new directory: `problems/challenges/22_message_decipherer/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
encodedMessage1 := "olleH ,dlroW"
encodedMessage2 := "oG si !gnizama"
encodedMessage3 := "év café 🚀" // A string with multi-byte characters

decodedMessage1 := decipherMessage(encodedMessage1)
fmt.Println(decodedMessage1)

decodedMessage2 := decipherMessage(encodedMessage2)
fmt.Println(decodedMessage2)

decodedMessage3 := decipherMessage(encodedMessage3)
fmt.Println(decodedMessage3)

// Expected Terminal Output:
Hello, World
Go is amazing!
vé café 🚀