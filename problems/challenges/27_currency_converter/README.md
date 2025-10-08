# Challenge #27: Currency Converter

## Objective
Write a Go program that simulates a currency converter. You will create a `Converter` that holds exchange rates and a method to convert an amount from one currency to another, parsing an input string.

## Difficulty
Medium

## Concepts Tested
* Maps (`map[string]float64`)
* Structs and Methods
* String Manipulation (`strings.Split`, `strings.ToUpper`)
* Error Handling
* String to Float Conversion (`strconv.ParseFloat`)

## Rules/Logic
1.  Define a struct named `Converter` with one field: `rates` (a map of type `map[string]float64`). The keys will be currency codes (e.g., "USD"), and the values will be their exchange rate relative to a base currency (e.g., USD).
2.  Create a `NewConverter()` constructor that returns an initialized `Converter` with some pre-defined exchange rates. For simplicity, we'll use USD as the base.
    * USD: 1.0
    * EUR: 0.93
    * INR: 83.30
3.  Implement a method `Convert(query string) (float64, error)`. The `query` string will be in the format `"AMOUNT FROM_CURRENCY TO_CURRENCY"` (e.g., `"100 USD INR"`).
4.  Inside the `Convert` method:
    * Split the query string into its three parts. If it doesn't have exactly three parts, return an error.
    * Parse the `AMOUNT` part from a string to a `float64`. Handle any parsing errors.
    * Convert the currency codes to uppercase to make the lookup case-insensitive.
    * Look up the exchange rates for both the `FROM_CURRENCY` and `TO_CURRENCY` in your `rates` map. If either currency is not found, return an error.
5.  **The Conversion Logic**: To convert from any currency `A` to `B`, the formula is: `(amount / rate_of_A) * rate_of_B`. This converts the amount to the base currency (USD) first, and then to the target currency.
6.  Return the final converted amount and a `nil` error.
7.  In your `main` function, test your `Convert` method with a few different queries, including one with an unknown currency and one with a malformed query string.

## Your Tasks
1.  Create a new directory: `problems/challenges/27_currency_converter/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your solution in `main.go`.

## Sample Input/Output

```go
// Sample code in your main function:
converter := NewConverter()

// Test case 1: USD to INR
result1, err1 := converter.Convert("100 USD INR")
if err1 != nil {
    fmt.Println("Error:", err1)
} else {
    fmt.Printf("100 USD is %.2f INR\n", result1)
}

// Test case 2: EUR to USD
result2, err2 := converter.Convert("50 EUR USD")
if err2 != nil {
    fmt.Println("Error:", err2)
} else {
    fmt.Printf("50 EUR is %.2f USD\n", result2)
}

// Test case 3: Unknown currency
_, err3 := converter.Convert("100 JPY INR")
if err3 != nil {
    fmt.Println("Error:", err3)
}

// Test case 4: Malformed query
_, err4 := converter.Convert("100 USD")
if err4 != nil {
    fmt.Println("Error:", err4)
}

// Expected Terminal Output:
100 USD is 8330.00 INR
50 EUR is 53.76 USD
Error: currency 'JPY' not found
Error: invalid query format. expected 'AMOUNT FROM_CURRENCY TO_CURRENCY'