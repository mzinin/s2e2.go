package s2e2

import "fmt"

// Converts infix token sequence into postfix one.
// Convertion is done by Shunting Yard algorithm.
type infixConverter struct {
	outputQueue   []token        // Output queue of all tokens.
	operatorStack []token        // Stack of operators and functions.
	operators     map[string]int // All expected operators and their priorities.
}

// Create new infix converter.
func newInfixConverter() *infixConverter {
	result := &infixConverter{}
	result.operators = make(map[string]int)
	return result
}

// Add operator expected within expression.
// Returns error if operator's name is not unique.
func (c *infixConverter) AddOperator(name string, priority int) error {
	if _, ok := c.operators[name]; ok {
		return fmt.Errorf("Converter: operator %v is already added", name)
	}

	c.operators[name] = priority
	return nil
}

// Convert infox token sequence into postfix one.
// Returns error if something goes wrong.
func (c *infixConverter) Convert(tokens []token) ([]token, error) {
	c.outputQueue = make([]token, 0)
	c.operatorStack = make([]token, 0)

	if err := c.processTokens(tokens); err != nil {
		return nil, err
	}
	if err := c.processOperators(); err != nil {
		return nil, err
	}

	return c.outputQueue, nil
}

// Process all tokens in the input sequence.
// Returns error if something goes wrong.
func (c *infixConverter) processTokens(tokens []token) error {
	for i := range tokens {
		switch tokens[i].Type {
		case atomType:
			c.processAtom(tokens[i])

		case commaType:
			c.processComma()

		case functionType:
			c.processFunction(tokens[i])

		case operatorType:
			if err := c.processOperator(tokens[i]); err != nil {
				return err
			}

		case leftBracketType:
			c.processLeftBracket(tokens[i])

		case rightBracketType:
			if err := c.processRightBracket(); err != nil {
				return err
			}

		default:
			return fmt.Errorf("Converter: unexpected token type %v", tokens[i].Type)
		}
	}

	return nil
}

// Process all operators left in the operator stack.
// Returns error if something goes wrong.
func (c *infixConverter) processOperators() error {
	for len(c.operatorStack) != 0 {
		last := len(c.operatorStack) - 1

		if c.operatorStack[last].Type == leftBracketType {
			return fmt.Errorf("Converter: unpaired bracket")
		}

		c.outputQueue = append(c.outputQueue, c.operatorStack[last])
		c.operatorStack = c.operatorStack[:last]
	}
	return nil
}

// Process ATOM token.
func (c *infixConverter) processAtom(tkn token) {
	c.outputQueue = append(c.outputQueue, tkn)
}

// Process COMMA token.
func (c *infixConverter) processComma() {
	for last := len(c.operatorStack) - 1; last >= 0 && c.operatorStack[last].Type != leftBracketType; last-- {
		c.outputQueue = append(c.outputQueue, c.operatorStack[last])
		c.operatorStack = c.operatorStack[:last]
	}
}

// Process FUNCTION token.
func (c *infixConverter) processFunction(tkn token) {
	c.operatorStack = append(c.operatorStack, tkn)
}

// Process OPERATOR token.
// Returns error in case of an unknown operator.
func (c *infixConverter) processOperator(tkn token) error {
	priority, ok := c.operators[tkn.Value]
	if !ok {
		return fmt.Errorf("Converter: unknown operator %v", tkn.Value)
	}

	for last := len(c.operatorStack) - 1; last >= 0; last-- {
		if c.operatorStack[last].Type != operatorType {
			break
		}
		if priority > c.operators[c.operatorStack[last].Value] {
			break
		}
		c.outputQueue = append(c.outputQueue, c.operatorStack[last])
		c.operatorStack = c.operatorStack[:last]
	}

	c.operatorStack = append(c.operatorStack, tkn)
	return nil
}

// Process LEFT BRACKET token.
func (c *infixConverter) processLeftBracket(tkn token) {
	c.operatorStack = append(c.operatorStack, tkn)
}

// Process RIGHT BRACKET token.
// Returns error in case of unpaired bracket.
func (c *infixConverter) processRightBracket() error {
	last := len(c.operatorStack) - 1

	for ; last >= 0; last-- {
		if c.operatorStack[last].Type == leftBracketType {
			break
		}
		c.outputQueue = append(c.outputQueue, c.operatorStack[last])
		c.operatorStack = c.operatorStack[:last]
	}

	if last < 0 {
		return fmt.Errorf("Converter: unpaired bracket")
	}
	c.operatorStack = c.operatorStack[:last]
	last--

	if last >= 0 && c.operatorStack[last].Type == functionType {
		c.outputQueue = append(c.outputQueue, c.operatorStack[last])
		c.operatorStack = c.operatorStack[:last]
	}
	return nil
}
