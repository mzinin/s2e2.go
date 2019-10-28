package s2e2

// Type of a token, kind of enum.
type tokenType int

const (
	// String literal, unsplittable.
	atomType tokenType = 0

	// Comma, used to separate function arguments.
	commaType tokenType = 1

	// Function, always followed by brackets with arguments.
	functionType tokenType = 2

	// Infix operator. Unlike functions does not use brackets,
	// can have arguments both before and after itself.
	operatorType tokenType = 3

	// Expression is either an atom or combination of several tokens.
	// Can be splitted.
	expressionType tokenType = 4

	// Left, opening round bracket.
	leftBracketType tokenType = 5

	// Right, closed round bracket.
	rightBracketType tokenType = 6
)
