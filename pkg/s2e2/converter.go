package s2e2

// Interface for converting token sequence.
type converter interface {
	// Add operator expected within expression.
	// Returns error if operator's name is not unique.
	AddOperator(name string, priority int) error

	// Convert token sequence.
	// Returns error if something goes wrong.
	Convert(tokens []token) ([]token, error)
}
