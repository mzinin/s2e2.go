package s2e2

import (
	"fmt"
	"sort"
	"strings"
)

// infixTokenizer splits infix string expressions into list of tokens.
type infixTokenizer struct {
	functions         map[string]bool         // Set of expected functions.
	operators         map[string]bool         // Set of expected operators.
	operatorsByLength map[int]map[string]bool // Operators sorted by their lengthes.
}

// newInfixTokenizer creates new infix tokenizer.
func newInfixTokenizer() *infixTokenizer {
	result := &infixTokenizer{}
	result.functions = make(map[string]bool)
	result.operators = make(map[string]bool)
	result.operatorsByLength = make(map[int]map[string]bool)
	return result
}

// AddFunction adds function expected within expression.
// Returns error if functons's name is not unique.
func (t *infixTokenizer) AddFunction(function string) error {
	if err := t.checkUniqueness(function); err != nil {
		return err
	}

	t.functions[function] = true
	return nil
}

// AddOperator adds operator expected within expression.
// Returns error if operator's name is not unique.
func (t *infixTokenizer) AddOperator(operator string) error {
	if err := t.checkUniqueness(operator); err != nil {
		return err
	}

	t.operators[operator] = true

	_, ok := t.operatorsByLength[len(operator)]
	if !ok {
		t.operatorsByLength[len(operator)] = make(map[string]bool)
	}
	t.operatorsByLength[len(operator)][operator] = true

	return nil
}

// Tokenize splits expression into tokens.
// Returns error if expression contains unknown symbol.
func (t *infixTokenizer) Tokenize(expression string) ([]token, error) {
	splitter, err := newExpressionSplitter(func(value string) tokenType {
		return t.tokenTypeByValue(value)
	})
	if err != nil {
		return nil, err
	}

	rawTokens, err := splitter.SplitIntoTokens(expression)
	if err != nil {
		return nil, err
	}

	refinedTokens := t.splitTokensByOperators(rawTokens)
	t.convertExpressionsIntoAtoms(refinedTokens)

	return refinedTokens, nil
}

// checkUniqueness checks if function's or operator's name is unique.
// Returns error if the name is not unique.
func (t *infixTokenizer) checkUniqueness(entityName string) error {
	if _, ok := t.functions[entityName]; ok {
		return fmt.Errorf("Tokenizer: function %v is already added", entityName)
	}
	if _, ok := t.operators[entityName]; ok {
		return fmt.Errorf("Tokenizer: operator %v is already added", entityName)
	}
	return nil
}

// tokenTypeByValue gets token type by its value.
func (t *infixTokenizer) tokenTypeByValue(value string) tokenType {
	if _, ok := t.operators[value]; ok {
		return operatorType
	}
	if _, ok := t.functions[value]; ok {
		return functionType
	}
	return expressionType
}

// splitTokensByOperators splits all tokens by all expected operatos.
func (t *infixTokenizer) splitTokensByOperators(tokens []token) []token {
	var lengths []int
	for length := range t.operatorsByLength {
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)

	result := tokens

	for _, length := range lengths {
		operatorsOfTheSameLength := t.operatorsByLength[length]
		for operatorName := range operatorsOfTheSameLength {
			result = t.splitTokensBySingleOperator(result, operatorName)
		}
	}

	return result
}

// splitTokensBySingleOperator splits all tokens by one operator.
func (t *infixTokenizer) splitTokensBySingleOperator(tokens []token, operator string) []token {
	result := make([]token, 0, len(tokens))

	for _, tkn := range tokens {
		if tkn.Type == expressionType {
			result = append(result, t.splitSingleTokenBySingleOperator(tkn.Value, operator)...)
		} else {
			result = append(result, tkn)
		}
	}

	return result
}

// splitSingleTokenBySingleOperator splits one token by one operator.
func (t *infixTokenizer) splitSingleTokenBySingleOperator(tokenValue string, operator string) []token {
	result := make([]token, 0)

	for start := 0; start < len(tokenValue); {
		if end := strings.Index(tokenValue[start:], operator); end == -1 {
			newValue := tokenValue[start:]
			newType := t.tokenTypeByValue(newValue)
			result = append(result, token{newType, newValue})
			break
		} else {
			if end != start {
				newValue := tokenValue[start:end]
				newType := t.tokenTypeByValue(newValue)
				result = append(result, token{newType, newValue})
			}
			newValue := tokenValue[end : end+len(operator)]
			result = append(result, token{operatorType, newValue})
			start = end + len(operator)
		}
	}

	return result
}

// convertExpressionsIntoAtoms converts all EXPRESSION tokens into ATOM ones.
func (t *infixTokenizer) convertExpressionsIntoAtoms(tokens []token) {
	for i := range tokens {
		if tokens[i].Type == expressionType {
			tokens[i].Type = atomType
		}
	}
}
