package s2e2

import "testing"

func TestInfixConverter_Positive_OneBinaryOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}}
	expectedTokens := []token{{atomType, "A"}, {atomType, "B"}, {operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_TwoBinaryOperatorsSamePriority_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := converter.AddOperator("-", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "-"}, {atomType, "C"}}
	expectedTokens := []token{{atomType, "A"}, {atomType, "B"}, {operatorType, "+"}, {atomType, "C"}, {operatorType, "-"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_TwoOperatorsDifferentPriorities_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := converter.AddOperator("*", 2); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "*"}, {atomType, "C"}}
	expectedTokens := []token{{atomType, "A"}, {atomType, "B"}, {atomType, "C"}, {operatorType, "*"}, {operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_UnaryOperatorAndBinaryOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("!=", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := converter.AddOperator("!", 2); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{{operatorType, "!"}, {atomType, "A"}, {operatorType, "!="}, {atomType, "B"}}
	expectedTokens := []token{{atomType, "A"}, {operatorType, "!"}, {atomType, "B"}, {operatorType, "!="}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_OneFunctionWithoutArguments_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {leftBracketType, "("}, {rightBracketType, ")"}}
	expectedTokens := []token{{functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_OneFunctionOneArgument_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {leftBracketType, "("}, {atomType, "Arg1"}, {rightBracketType, ")"}}
	expectedTokens := []token{{atomType, "Arg1"}, {functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_OneFunctionThreeArguments_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{
		{functionType, "FUN"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{atomType, "Arg2"},
		{atomType, "Arg3"},
		{rightBracketType, ")"}}

	expectedTokens := []token{{atomType, "Arg1"}, {atomType, "Arg2"}, {atomType, "Arg3"}, {functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_FunctionAndExernalOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{
		{functionType, "FUN"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{rightBracketType, ")"},
		{operatorType, "+"},
		{functionType, "FUN"},
		{leftBracketType, "("},
		{atomType, "Arg2"},
		{rightBracketType, ")"}}

	expectedTokens := []token{
		{atomType, "Arg1"},
		{functionType, "FUN"},
		{atomType, "Arg2"},
		{functionType, "FUN"},
		{operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_FunctionAndInternalOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{
		{functionType, "FUN"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{operatorType, "+"},
		{atomType, "Arg2"},
		{commaType, ","},
		{atomType, "Arg3"},
		{operatorType, "+"},
		{atomType, "Arg4"},
		{rightBracketType, ")"}}

	expectedTokens := []token{
		{atomType, "Arg1"},
		{atomType, "Arg2"},
		{operatorType, "+"},
		{atomType, "Arg3"},
		{atomType, "Arg4"},
		{operatorType, "+"},
		{functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_NestedFunctions_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{
		{functionType, "FUN1"},
		{leftBracketType, "("},
		{functionType, "FUN2"},
		{leftBracketType, "("},
		{rightBracketType, ")"},
		{commaType, ","},
		{functionType, "FUN3"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{commaType, ","},
		{atomType, "Arg2"},
		{rightBracketType, ")"},
		{rightBracketType, ")"}}

	expectedTokens := []token{
		{functionType, "FUN2"},
		{atomType, "Arg1"},
		{atomType, "Arg2"},
		{functionType, "FUN3"},
		{functionType, "FUN1"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_OperatorsWithoutArguments_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{{operatorType, "+"}, {operatorType, "+"}, {operatorType, "+"}}
	expectedTokens := []token{{operatorType, "+"}, {operatorType, "+"}, {operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_FunctionWithoutCommas_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{
		{functionType, "FUN"},
		{leftBracketType, "("},
		{atomType, "Arg1"},
		{atomType, "Arg2"},
		{rightBracketType, ")"}}

	expectedTokens := []token{{atomType, "Arg1"}, {atomType, "Arg2"}, {functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Positive_FunctionOfOperators_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{
		{functionType, "FUN"},
		{leftBracketType, "("},
		{operatorType, "+"},
		{operatorType, "+"},
		{rightBracketType, ")"}}

	expectedTokens := []token{{operatorType, "+"}, {operatorType, "+"}, {functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	if err != nil {
		test.Errorf("Unexpected error: %v", err)
		return
	}

	compareTokens(actualTokens, expectedTokens, test)
}

func TestInfixConverter_Negative_UnpairedLeftBracket(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {leftBracketType, "("}, {atomType, "Arg1"}}

	if _, err := converter.Convert(inputTokens); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Converter: unpaired bracket" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestInfixConverter_Negative_UnpairedRightBracket(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {atomType, "Arg1"}, {rightBracketType, ")"}}

	if _, err := converter.Convert(inputTokens); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Converter: unpaired bracket" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestInfixConverter_Negative_TwoOperatorsWithTheSameName(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := converter.AddOperator("+", 1); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Converter: operator + is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestInfixConverter_Negative_UnknownOperator(test *testing.T) {
	converter := newInfixConverter()

	if err := converter.AddOperator("+", 1); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	inputTokens := []token{
		{atomType, "Arg1"},
		{operatorType, "+"},
		{atomType, "Arg2"},
		{operatorType, "*"},
		{atomType, "Arg3"}}

	if _, err := converter.Convert(inputTokens); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Converter: unknown operator *" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
