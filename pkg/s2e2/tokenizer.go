package s2e2

// tokenizer is the interface for splitting expression string into list of tokens.
type tokenizer interface {

	// AddFunction adds function expected within expression.
	// Returns error if functons's name is not unique.
	AddFunction(function string) error

	// AddOperator adds operator expected within expression.
	// Returns error if operator's name is not unique.
	AddOperator(operator string) error

	// Tokenize splits expression into tokens.
	// Returns error if expression contains unknown symbol.
	Tokenize(expression string) ([]token, error)
}
