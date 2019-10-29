package s2e2

import (
	"fmt"
	"strings"
	"unicode"
)

// Some special symbols.
const (
	commaSymbol        = ','
	leftBracketSymbol  = '('
	rightBracketSymbol = ')'
	quoteSymbol        = '"'
	backslashSymbol    = '\\'
)

// Splits expression into raw tokens.
type expressionSplitter struct {
	insideQuotes bool                   // Flag of "inside quotes" state. If set it means that current symbol belongs to an ATOM.
	currentToken string                 // Currently parsed token value.
	foundTokens  []token                // List of found tokens.
	typeByValue  func(string) tokenType // External function to get token's type by its value.
}

// Create new expression splitter. Returns error if the provided external function is nil.
func newExpressionSplitter(typeByValue func(string) tokenType) (*expressionSplitter, error) {
	if typeByValue == nil {
		return nil, fmt.Errorf("Splitter: external function to get token type by its value is nil")
	}

	result := &expressionSplitter{}
	result.insideQuotes = false
	result.currentToken = ""
	result.foundTokens = make([]token, 0)
	result.typeByValue = typeByValue
	return result, nil
}

// Split expression into tokens by spaces and brackets.
// Returns error if expression contains unknown symbol or external function is nil.
func (s *expressionSplitter) SplitIntoTokens(expression string) ([]token, error) {
	for _, symbol := range expression {
		if err := s.processSymbol(symbol); err != nil {
			return nil, err
		}
	}
	if err := s.flushToken(); err != nil {
		return nil, err
	}
	return s.foundTokens, nil
}

// Process one symbol of the input expression.
// Returns error if expression contains unknown symbol or external function is nil.
func (s *expressionSplitter) processSymbol(symbol rune) error {
	switch symbol {
	case commaSymbol, leftBracketSymbol, rightBracketSymbol:
		return s.processSpecialSymbol(symbol)

	case quoteSymbol:
		return s.processQuoteSymbol(symbol)

	default:
		return s.processCommonSymbol(symbol)
	}
}

// Process one special symbol of the input expression.
// Returns error if expression contains unknown symbol or external function is nil.
func (s *expressionSplitter) processSpecialSymbol(symbol rune) error {
	if s.insideQuotes {
		s.addSymbolToToken(symbol)
		return nil
	}

	if err := s.flushToken(); err != nil {
		return err
	}

	switch symbol {
	case commaSymbol:
		s.addFoundToken(commaType, string(commaSymbol))

	case leftBracketSymbol:
		s.addFoundToken(leftBracketType, string(leftBracketSymbol))

	case rightBracketSymbol:
		s.addFoundToken(rightBracketType, string(rightBracketSymbol))

	default:
		return fmt.Errorf("Splitter: unexpected special symbol %v", symbol)
	}

	return nil
}

// Process one quote symbol of the input expression.
// Returns error if external function is nil.
func (s *expressionSplitter) processQuoteSymbol(symbol rune) error {
	if s.insideQuotes && s.isEscaped() {
		s.addSymbolToToken(symbol)
	}

	if err := s.flushToken(); err != nil {
		return err
	}

	s.insideQuotes = !s.insideQuotes
	return nil
}

// Process one common symbol of the input expression.
// Returns error if external function is nil.
func (s *expressionSplitter) processCommonSymbol(symbol rune) error {
	if s.insideQuotes || !unicode.IsSpace(symbol) {
		s.addSymbolToToken(symbol)
	} else if err := s.flushToken(); err != nil {
		return err
	}
	return nil
}

// Add symbol to currently parsed token.
func (s *expressionSplitter) addSymbolToToken(symbol rune) {
	if symbol == quoteSymbol {
		s.currentToken = s.currentToken[:len(s.currentToken)-1] + string(symbol)
	} else {
		s.currentToken += string(symbol)
	}
}

// Add current token if there is such to the list of found tokens.
// Returns error if external function is nil.
func (s *expressionSplitter) flushToken() error {
	if !s.insideQuotes {
		s.currentToken = strings.TrimSpace(s.currentToken)
	}

	if len(s.currentToken) != 0 || s.insideQuotes {
		currentTokenType, err := s.tokenTypeByValue(s.currentToken)
		if err != nil {
			return err
		}
		s.addFoundToken(currentTokenType, s.currentToken)
	}

	s.currentToken = ""
	return nil
}

// Add token to the list of found tokens.
func (s *expressionSplitter) addFoundToken(typeOfToken tokenType, valueOfToken string) {
	s.foundTokens = append(s.foundTokens, token{typeOfToken, valueOfToken})
}

// Check if current symbol is escaped, i.e. preceded by a backslash.
func (s *expressionSplitter) isEscaped() bool {
	return len(s.currentToken) != 0 && s.currentToken[len(s.currentToken)-1] == backslashSymbol
}

// Get token type by its value and current state of the splitter.
// Returns error if external function is nil.
func (s *expressionSplitter) tokenTypeByValue(value string) (tokenType, error) {
	if s.insideQuotes {
		return atomType, nil
	}
	if s.typeByValue == nil {
		return atomType, fmt.Errorf("Splitter: external function to get token type by its value is nil")
	}
	return s.typeByValue(value), nil
}
