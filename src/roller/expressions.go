package roller

import (
	"math/rand"
	"strconv"
)

type Expression interface {
	Evaluate() int64
}

type BinaryExpression struct {
	Left  Expression
	Op    string
	Right Expression
}

type UnaryExpression struct {
	Op   string
	Expr Expression
}

type ConstantExpression struct {
	Value int64
}

func Parse(tokens []Token) Expression {
	// Implement parsing logic here
	index := 0
	var currentExpression Expression
	for {
		token := tokens[index]
		switch token.Kind {
		case "number":
			// Handle number token
			n, _ := strconv.ParseInt(token.Value, 10, 32)
			currentExpression = &ConstantExpression{Value: n} // Convert char to int
			index++
		case "dice":
			// Handle dice token
			left := currentExpression
			index++
			right := Parse(tokens[index:]) // Recursively parse the right side
			return Roll(left, right)
		case "+":
			// Handle addition token
			left := currentExpression
			index++
			right := Parse(tokens[index:]) // Recursively parse the right side
			return Add(left, right)
		case "-":
			// Handle subtraction token
			left := currentExpression
			index++
			right := Parse(tokens[index:]) // Recursively parse the right side
			return Subtract(left, right)
		default:
			panic("Unexpected token: " + token.Kind)
		}
		if index >= len(tokens) {
			break
		}
	}

	return currentExpression
}

func Roll(left Expression, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Left:  left,
		Op:    "d",
		Right: right,
	}
}

func Add(left Expression, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Left:  left,
		Op:    "+",
		Right: right,
	}
}

func Subtract(left Expression, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Left:  left,
		Op:    "-",
		Right: right,
	}
}

func Negate(expr Expression) *UnaryExpression {
	return &UnaryExpression{
		Op:   "-",
		Expr: expr,
	}
}

func (expr *ConstantExpression) Evaluate() int64 {
	return expr.Value
}

func (expr *BinaryExpression) Evaluate() int64 {
	leftResult := expr.Left.Evaluate()
	rightResult := expr.Right.Evaluate()

	switch expr.Op {
	case "d":
		// Implement dice rolling logic here
		sum := int64(0)
		for i := int64(0); i < leftResult; i++ {
			// Roll a die with rightResult sides and accumulate the total
			sum += int64(1 + (rand.Intn(int(rightResult))))
		}
		return sum
	case "+":
		return leftResult + rightResult
	case "-":
		return leftResult - rightResult
	default:
		panic("Unknown operator")
	}
}
