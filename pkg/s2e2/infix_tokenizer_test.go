package s2e2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfixTokenizer_Positive_OneOperatorWithSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))

	expression := "A + B"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_OneOperatorWithoutSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))

	expression := "A+B"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_TwoOperatorWithSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))
	assert.NoError(test, tokenizer.AddOperator("&&"))

	expression := "A + B && C"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "&&"}, {atomType, "C"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_TwoOperatorWithoutSpaces_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))
	assert.NoError(test, tokenizer.AddOperator("&&"))

	expression := "A+B&&C"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "+"}, {atomType, "B"}, {operatorType, "&&"}, {atomType, "C"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_OneOperatorIsSubstringOfAnother_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("!"))
	assert.NoError(test, tokenizer.AddOperator("!="))

	expression := "A != !B"
	expectedTokens := []token{{atomType, "A"}, {operatorType, "!="}, {operatorType, "!"}, {atomType, "B"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_OneFunctionWithoutArguments_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FUN1"))

	expression := "FUN1()"
	expectedTokens := []token{{functionType, "FUN1"}, {leftBracketType, "("}, {rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_OneFunctionOneArgument_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FUN1"))

	expression := "FUN1(Arg1)"
	expectedTokens := []token{{functionType, "FUN1"}, {leftBracketType, "("}, {atomType, "Arg1"}, {rightBracketType, ")"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_OneFunctionThreeArguments_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FUN1"))

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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_TwoFunctionsOneOperator_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FUN1"))
	assert.NoError(test, tokenizer.AddFunction("FUN2"))
	assert.NoError(test, tokenizer.AddOperator("+"))

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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_NestedFunctions_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FUN1"))
	assert.NoError(test, tokenizer.AddFunction("FUN2"))
	assert.NoError(test, tokenizer.AddFunction("FUN3"))

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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_NestedBrackets_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))

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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Positive_OperatorsWithoutArguments_ResultValue(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))

	expression := "+ + +"
	expectedTokens := []token{
		{operatorType, "+"},
		{operatorType, "+"},
		{operatorType, "+"}}

	actualTokens, err := tokenizer.Tokenize(expression)
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
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
	assert.NoError(test, err)
	assert.Equal(test, expectedTokens, actualTokens)
}

func TestInfixTokenizer_Negative_TwoOperatorsWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("+"))

	err := tokenizer.AddOperator("+")
	assert.Error(test, err)
	assert.Equal(test, "Tokenizer: operator + is already added", err.Error())
}

func TestInfixTokenizer_Negative_TwoFunctionsWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FUN"))

	err := tokenizer.AddFunction("FUN")
	assert.Error(test, err)
	assert.Equal(test, "Tokenizer: function FUN is already added", err.Error())
}

func TestInfixTokenizer_Negative_FunctionAndOperatorWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddFunction("FF"))

	err := tokenizer.AddOperator("FF")
	assert.Error(test, err)
	assert.Equal(test, "Tokenizer: function FF is already added", err.Error())
}

func TestInfixTokenizer_Negative_OperatorAndFunctionWithTheSameName(test *testing.T) {
	tokenizer := newInfixTokenizer()

	assert.NoError(test, tokenizer.AddOperator("FF"))

	err := tokenizer.AddFunction("FF")
	assert.Error(test, err)
	assert.Equal(test, "Tokenizer: operator FF is already added", err.Error())
}
