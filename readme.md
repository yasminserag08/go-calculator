# CLI Calculator in Go

A command-line calculator written in Go. It handles operator precedence and nested parentheses, tracks calculation history, and includes error handling.

---

## Features

* **Expression Evaluation:** Supports addition (`+`), subtraction (`-`), multiplication (`*`), division (`/`), and nested parentheses `( )`.
* **Session History:** Automatically logs past calculations, allowing you to view or clear your history at any time.
* **Error Handling:** Catches malformed expressions, division by zero, and mismatched parentheses gracefully without crashing.

---

## Project Structure

* `calculator.go` — Contains the core mathematical parsing engine, infix to postfix conversion, postfix evaluation, and operation logic.
* `history.go` — Manages storage, viewing, and clearing of past calculation sessions.
* `main.go` — Handles the main application control flow and user input loop.
* `menu.go` — Handles menu rendering and user choice validation.

---

## How to Run

Run the program using:

```bash
go run .