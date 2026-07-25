# CLI Calculator in Go

A command-line calculator written in Go. It handles operator precedence and nested parentheses, tracks calculation history, and includes error handling.

---

## Features

* **Expression Evaluation:** Supports addition (`+`), subtraction (`-`), multiplication (`*`), division (`/`), and nested parentheses `( )`.
* **Session History:** Automatically logs past calculations, allowing you to view or clear your history at any time.
* **Error Handling:** Catches malformed expressions, division by zero, and mismatched parentheses gracefully without crashing.

---

## Project Structure

* `calculator.go` — Contains infix to postfix conversion, postfix evaluation, and operation logic.
* `history.go` — Manages storage, viewing, and clearing of past calculation sessions.
* `menu.go` — Handles menu rendering and user choice validation.
* `parsing.go` — Manages tokenization and expression normalization.
* `stack.go` — Contains all generic stack operations.
* `main.go` — Handles the main application control flow and user input loop.

---

## How to Run

Run the program using:

```bash
go run .