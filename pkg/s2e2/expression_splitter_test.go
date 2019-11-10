package s2e2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionSplitter_Positive_NewSplitterResult(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	_, err := newExpressionSplitter(typeByValue)
	assert.NoError(test, err)
}

func TestExpressionSplitter_Positive_SplitByComma(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	assert.NoError(test, err)

	expression := "A, B"
	expectedTokens := []token{{atomType, "A"}, {commaType, ","}, {atomType, "B"}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestExpressionSplitter_Positive_SplitByBrackets(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	assert.NoError(test, err)

	expression := "(A, B)"
	expectedTokens := []token{
		{leftBracketType, "("},
		{atomType, "A"},
		{commaType, ","},
		{atomType, "B"},
		{rightBracketType, ")"}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestExpressionSplitter_Positive_QuotedAtom(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	assert.NoError(test, err)

	expression := "A, \"B C\""
	expectedTokens := []token{{atomType, "A"}, {commaType, ","}, {atomType, "B C"}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestExpressionSplitter_Positive_QuotedUntrimmedAtom(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	assert.NoError(test, err)

	expression := "A, \" B \""
	expectedTokens := []token{{atomType, "A"}, {commaType, ","}, {atomType, " B "}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestExpressionSplitter_Negative_NilExternalFunction(test *testing.T) {
	_, err := newExpressionSplitter(nil)
	assert.Error(test, err)
	assert.Equal(test, "Splitter: external function to get token type by its value is nil", err.Error())
}
