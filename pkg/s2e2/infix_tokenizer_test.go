package s2e2

import "testing"

func TestInfixTokenizer_Positive_OneOperatorWithSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "A + B"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_OneOperatorWithoutSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "A+B"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_TwoOperatorWithSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddOperator("&&"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "A + B && C"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "&&"}, {atomType, "C"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_TwoOperatorWithoutSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddOperator("&&"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "A+B&&C"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "&&"}, {atomType, "C"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_OneOperatorIsSubstringOfAnother_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("!"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddOperator("!="); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "A != !B"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "!="}, {operatorType, "!"}, {atomType, "B"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_OneFunctionWithoutArguments_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FUN1"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "FUN1()"
	expectedTokens := []token{{functionType, "FUN1"}, {leftBracketType, "("}, {rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_OneFunctionOneArgument_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FUN1"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "FUN1(Arg1)"
	expectedTokens := []token{{functionType, "FUN1"}, {leftBracketType, "("}, {atomType, "Arg1"}, {rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_OneFunctionThreeArguments_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FUN1"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "FUN1(Arg1, Arg2,Arg3)"
	expectedTokens := []token{
		{functionType, "FUN1"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{commaType, ","},
		{atomType, "Arg2"},
		{commaType, ","},
		{atomType, "Arg3"},
		{rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_TwoFunctionsOneOperator_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FUN1"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddFunction("FUN2"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "FUN1(Arg1) + FUN2(Arg2)"
	expectedTokens := []token{
		{functionType, "FUN1"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{rightBracketType, ")"},
		{operatorType, "+"},
		{functionType, "FUN2"},
		{leftBracketType, "("},
		{atomType, "Arg2"},
		{rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_NestedFunctions_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FUN1"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddFunction("FUN2"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := tokenizer.AddFunction("FUN3"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "FUN1(FUN2(), FUN3())"
	expectedTokens := []token{
		{functionType, "FUN1"},
		{leftBracketType, "("},
		{functionType, "FUN2"},
		{leftBracketType, "("},
		{rightBracketType, ")"},
		{commaType, ","},
		{functionType, "FUN3"},
		{leftBracketType, "("},
		{rightBracketType, ")"},
		{rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_NestedBrackets_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "(((A + B)))"
	expectedTokens := []token{
		{leftBracketType, "("},
		{leftBracketType, "("},
		{leftBracketType, "("},
		{atomType, "A"},
		{operatorType, "+"},
		{atomType, "B"},
		{rightBracketType, ")"},
		{rightBracketType, ")"},
		{rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_OperatorsWithoutArguments_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	expression := "+ + +"
	expectedTokens := []token{
		{operatorType, "+"},
		{operatorType, "+"},
		{operatorType, "+"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Positive_UnpairedBrackets_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	expression := "((()"
	expectedTokens := []token{
		{leftBracketType, "("},
		{leftBracketType, "("},
		{leftBracketType, "("},
		{rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixTokenizer_Negative_TwoOperatorsWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("+"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := tokenizer.AddOperator("+"); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Tokenizer: operator + is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestInfixTokenizer_Negative_TwoFunctionsWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FUN"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := tokenizer.AddFunction("FUN"); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Tokenizer: function FUN is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestInfixTokenizer_Negative_FunctionAndOperatorWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddFunction("FF"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := tokenizer.AddOperator("FF"); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Tokenizer: function FF is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestInfixTokenizer_Negative_OperatorAndFunctionWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	if err := tokenizer.AddOperator("FF"); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := tokenizer.AddFunction("FF"); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Tokenizer: operator FF is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
