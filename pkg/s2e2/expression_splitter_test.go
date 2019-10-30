package s2e2

import "testing"

func TestExpressionSplitter_Positive_NewSplitterResult(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	if _, err := newExpressionSplitter(typeByValue); err != nil {
		test.Errorf("Unexpected error: %v", err)
	}
}

func TestExpressionSplitter_Positive_SplitByComma(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	expression := "A, B"
	expectedTokens := []token{{atomType, "A"}, {commaType, ","}, {atomType, "B"}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestExpressionSplitter_Positive_SplitByBrackets(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	expression := "(A, B)"
	expectedTokens := []token{
		{leftBracketType, "("},
		{atomType, "A"},
		{commaType, ","},
		{atomType, "B"},
		{rightBracketType, ")"}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestExpressionSplitter_Positive_QuotedAtom(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	expression := "A, \"B C\""
	expectedTokens := []token{{atomType, "A"}, {commaType, ","}, {atomType, "B C"}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestExpressionSplitter_Positive_QuotedUntrimmedAtom(test *testing.T) {
	typeByValue := func(string) tokenType {
		return atomType
	}

	splitter, err := newExpressionSplitter(typeByValue)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	expression := "A, \" B \""
	expectedTokens := []token{{atomType, "A"}, {commaType, ","}, {atomType, " B "}}

	actualTokens, err := splitter.SplitIntoTokens(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestExpressionSplitter_Negative_NilExternalFunction(test *testing.T) {
	if _, err := newExpressionSplitter(nil); err == nil {
		test.Errorf("No expected error")
		return
	} else if err.Error() != "Splitter: external function to get token type by its value is nil" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}