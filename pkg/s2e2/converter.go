package s2e2

// converter is the interface for converting token sequence.
type converter interface {

	// AddOperator adds operator expected within expression.
	// Returns error if operator's name is not unique.
	AddOperator(name string, priority int) error

	// Convert converts token sequence.
	// Returns error if something goes wrong.
	Convert(tokens []token) ([]token, error)
}
