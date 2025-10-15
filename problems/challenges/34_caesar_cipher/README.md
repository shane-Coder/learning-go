# Challenge #34: Caesar Cipher

## Objective
Write a Go program to implement a Caesar cipher, a simple method of encryption. The function will take a plaintext string and a shift key, and it will return the encrypted ciphertext.

## Difficulty
Medium

## Concepts Tested
* Strings and Runes
* Looping and Conditionals
* The Modulo Operator (`%`)
* Character Encoding (ASCII/Unicode)

## Rules/Logic
1.  Create a function `caesarCipher(text string, shift int) string`.
2.  The function should shift each letter in the `text` by the number of positions specified by `shift`.
3.  The cipher should only affect alphabetic characters (`a-z`, `A-Z`). All other characters (numbers, symbols, spaces) should remain unchanged.
4.  The shifting must "wrap around" the alphabet. For example, if you shift `'y'` by `4`, it should wrap around to `'c'`. The modulo operator (`% 26`) is perfect for this.
5.  The function must preserve the case of the original letters. Uppercase letters should remain uppercase, and lowercase letters should remain lowercase.
6.  **Hint**: It's easiest to work with runes. You can check if a rune is within the `'a'` to `'z'` range or `'A'` to `'Z'` range. To perform the shift, you can convert the rune to a 0-25 integer, apply the shift and the modulo, and then convert it back to a rune.
    * `(rune - 'a' + shift) % 26 + 'a'` for lowercase.
    * `(rune - 'A' + shift) % 26 + 'A'` for uppercase.

## Your Tasks
1.  Create a new directory: `problems/challenges/34_caesar_cipher/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
text1 := "Hello, World!"
shift1 := 3
encrypted1 := caesarCipher(text1, shift1)
fmt.Printf("Original: '%s'\nEncrypted: '%s'\n\n", text1, encrypted1)

text2 := "Zebra-123"
shift2 := 5
encrypted2 := caesarCipher(text2, shift2)
fmt.Printf("Original: '%s'\nEncrypted: '%s'\n", text2, encrypted2)

// Expected Terminal Output:
Original: 'Hello, World!'
Encrypted: 'Khoor, Zruog!'

Original: 'Zebra-123'
Encrypted: 'Ejgwf-123'