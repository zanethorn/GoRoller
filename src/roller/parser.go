package roller

type Token struct {
	Kind  string
	Value string
}

func Tokenize(input string) []Token {
	index := 0
	tokens := make([]Token, 0)

	for {
		k := input[index]
		switch k {
		case ' ':
			// Skip whitespace
		case '+':
			// Handle operator
			newToken := Token{Kind: "+", Value: "+"}
			tokens = append(tokens, newToken)
		case '-':
			newToken := Token{Kind: "-", Value: "-"}
			tokens = append(tokens, newToken)
		case '(':
			// Handle parentheses
			newToken := Token{Kind: "(", Value: "("}
			tokens = append(tokens, newToken)
		case ')':
			// Handle parentheses
			newToken := Token{Kind: ")", Value: ")"}
			tokens = append(tokens, newToken)
		case 'd':
			// Handle dice notation
			newToken := Token{Kind: "dice", Value: "d"}
			tokens = append(tokens, newToken)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// Handle number
			newToken := Token{Kind: "number", Value: string(k)}
			tokens = append(tokens, newToken)
		default:
			panic("Unexpected character: " + string(k))

		}
		if index >= len(input)-1 {
			break
		}
		index++
	}

	return tokens
}
