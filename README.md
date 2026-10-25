# My Monkey Programming Language in Go

It's my custom idea of the **Monkey programming language** in Go, inspired by the book [*Writing An Interpreter In Go*](https://interpreterbook.com/) by Thorsten Ball. 

This project features a complete lexer, parser, abstract syntax tree (AST), environment system, and a tree-walking evaluator capable of handling arithmetic expressions, bindings (`let`) and conditionals (`if-else`).

## Features

- **Lexer**: Tokenizes source code into meaningful tokens.
- **Parser**: Pratt parser building a robust Abstract Syntax Tree (AST).
- **Evaluator**: Tree-walking evaluator supporting integers, booleans, prefix/infix operations, conditionals, variable scoping, and function execution.
- **REPL**: Interactive Read-Eval-Print Loop to run code on the fly.

## Getting Started

Make sure you have [Go](https://golang.org/) installed on your machine.

### Clone the repository
```bash
git clone https://github.com/citunek22/monkey-interpreter.git
cd monkey-interpreter
```

### Run the tests
You can run the test suite across all packages using:
```bash
go test ./...
```

### Start the REPL
To launch the interactive console and start writing code, run:
```bash
go run cmd/monkey/main.go
```

## Example Usage

Once inside the REPL, you can evaluate expressions, assign variables, or run functions:

```monkey
>> let add = fn(a, b) { a + b; };
>> add(10, 15);
25
>> let x = 5 * 2;
>> x > 3;
true
```

## Acknowledgments / Credits (or how you call it)
Huge thanks to Thorsten Ball for his amazing book [*Writing An Interpreter In Go*](https://interpreterbook.com/), which guided the architecture of this project.

I also hope you don't mind that some parts are in czech because czech is my native language
