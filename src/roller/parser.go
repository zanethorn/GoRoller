package roller

type Token struct {
	Kind  string
	Value string
}

func Tokenize(input string) []Token {
	start := 0
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
			// Handle multi-character numbers
			start = index
			for index < len(input) && input[index] >= '0' && input[index] <= '9' {
				index++
			}
			newToken := Token{Kind: "number", Value: input[start:index]}
			tokens = append(tokens, newToken)
			index-- // Decrement because the main loop will increment
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
