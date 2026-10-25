package main

import (
	"bufio"
	"fmt"
	"monkey/internal/evaluator"
	"monkey/internal/lexer"
	"monkey/internal/object"
	"monkey/internal/parser"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ahoj %s! Toto je programovací jazyk Monkey!\n", currentUser.Username)
	fmt.Printf("Zadej nějaký příkaz:\n")

	scanner := bufio.NewScanner(os.Stdin)
	env := object.NewEnvironment()

	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Println("parser error: " + msg)
			}
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Println(evaluated.Inspect())
		}
	}
}
