package s2e2

// token i.e. a unit of some expression.
type token struct {
	Type  tokenType // Token type.
	Value string    // Token string value.
}
