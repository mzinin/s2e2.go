package s2e2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfixConverter_Positive_OneBinaryOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

	inputTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}}
	expectedTokens := []token{{atomType, "A"}, {atomType, "B"}, {operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_TwoBinaryOperatorsSamePriority_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))
	assert.NoError(test, converter.AddOperator("-", 1))

	inputTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "-"}, {atomType, "C"}}
	expectedTokens := []token{{atomType, "A"}, {atomType, "B"}, {operatorType, "+"}, {atomType, "C"}, {operatorType, "-"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_TwoOperatorsDifferentPriorities_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))
	assert.NoError(test, converter.AddOperator("*", 2))

	inputTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "*"}, {atomType, "C"}}
	expectedTokens := []token{{atomType, "A"}, {atomType, "B"}, {atomType, "C"}, {operatorType, "*"}, {operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_UnaryOperatorAndBinaryOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("!=", 1))
	assert.NoError(test, converter.AddOperator("!", 2))

	inputTokens := []token{{operatorType, "!"}, {atomType, "A"}, {operatorType, "!="}, {atomType, "B"}}
	expectedTokens := []token{{atomType, "A"}, {operatorType, "!"}, {atomType, "B"}, {operatorType, "!="}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_OneFunctionWithoutArguments_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {leftBracketType, "("}, {rightBracketType, ")"}}
	expectedTokens := []token{{functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_OneFunctionOneArgument_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {leftBracketType, "("}, {atomType, "Arg1"}, {rightBracketType, ")"}}
	expectedTokens := []token{{atomType, "Arg1"}, {functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_FunctionAndExernalOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_FunctionAndInternalOperator_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_OperatorsWithoutArguments_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

	inputTokens := []token{{operatorType, "+"}, {operatorType, "+"}, {operatorType, "+"}}
	expectedTokens := []token{{operatorType, "+"}, {operatorType, "+"}, {operatorType, "+"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Positive_FunctionOfOperators_ResultValue(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

	inputTokens := []token{
		{functionType, "FUN"},
		{leftBracketType, "("},
		{operatorType, "+"},
		{operatorType, "+"},
		{rightBracketType, ")"}}

	expectedTokens := []token{{operatorType, "+"}, {operatorType, "+"}, {functionType, "FUN"}}

	actualTokens, err := converter.Convert(inputTokens)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixConverter_Negative_UnpairedLeftBracket(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {leftBracketType, "("}, {atomType, "Arg1"}}

	_, err := converter.Convert(inputTokens)
	assert.Error(test, err)
	assert.Equal(test, "Converter: unpaired bracket", err.Error())
}

func TestInfixConverter_Negative_UnpairedRightBracket(test *testing.T) {
	converter := newInfixConverter()

	inputTokens := []token{{functionType, "FUN"}, {atomType, "Arg1"}, {rightBracketType, ")"}}

	_, err := converter.Convert(inputTokens)
	assert.Error(test, err)
	assert.Equal(test, "Converter: unpaired bracket", err.Error())
}

func TestInfixConverter_Negative_TwoOperatorsWithTheSameName(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

	err := converter.AddOperator("+", 1)
	assert.Error(test, err)
	assert.Equal(test, "Converter: operator + is already added", err.Error())
}

func TestInfixConverter_Negative_UnknownOperator(test *testing.T) {
	converter := newInfixConverter()

	assert.NoError(test, converter.AddOperator("+", 1))

	inputTokens := []token{
		{atomType, "Arg1"},
		{operatorType, "+"},
		{atomType, "Arg2"},
		{operatorType, "*"},
		{atomType, "Arg3"}}

	_, err := converter.Convert(inputTokens)
	assert.Error(test, err)
	assert.Equal(test, "Converter: unknown operator *", err.Error())
}
