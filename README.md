# Learning Go: A Daily Practice Repository 🚀

## What is this Repository?

This repository documents my daily journey of learning the Go programming language from the fundamentals to more advanced topics. My goal is not just to learn Go's syntax, but to build strong, professional habits in problem-solving, code structuring, and version control using Git.

Every day, I tackle new concepts in a structured lesson and then apply that knowledge by solving a practical problem.

## 📂 Project Structure

A key concept in Go is that a single executable program (defined as `package main`) can only have one `main()` function, which serves as its entry point.

To keep multiple, runnable examples and problems in one repository without causing a `'main' redeclared` error, each lesson and problem is organized into its own self-contained directory.

### `/lessons`
This directory contains the code examples from my daily learning sessions, organized chronologically. Each sub-folder is a standalone Go program demonstrating a specific concept.

### `/problems`
This directory contains my solutions to the daily problem-solving challenges. Each sub-folder is a standalone Go program that solves a specific problem, complete with the problem statement and sample output in the comments.

## 💻 How to Run an Example

To run any lesson or problem:
1.  Navigate into the specific directory.
    ```bash
    # Example for a lesson
    cd lessons/01_variables

    # Example for a problem
    cd problems/01_shipping_calculator
    ```
2.  Use the `go run` command.
    ```bash
    go run main.go
    ```

## ✨ Our Learning Framework

This journey is guided by a 7-point framework to ensure a structured and professional approach:

1.  **Guiding Principle**: Stay Current & Relevant
2.  **Setup First**: Create a Stable Environment
3.  **The Learning Loop**: Notes & Interview Prep
4.  **The Practical Loop**: Daily Problem-Solving
5.  **Professional Habits**: Use Git & Read Docs
6.  **The Review Cycle**: Daily & Weekly Summaries
7.  **Our Shared Process**: Reminders & Tracking