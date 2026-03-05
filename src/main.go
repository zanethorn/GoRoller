package main

import (
	"bufio"
	"fmt"
	"os"

	"main/roller"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	tokens := roller.Tokenize(input[:len(input)-1]) // Remove the newline character
	expressions := roller.Parse(tokens)

	result := expressions.Evaluate()
	fmt.Println(result)
}
