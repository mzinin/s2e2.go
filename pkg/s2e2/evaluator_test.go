package s2e2

import (
	"fmt"

	"github.com/mzinin/s2e2.go/pkg/s2e2/functions"
	"github.com/mzinin/s2e2.go/pkg/s2e2/operators"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestEvaluator_Positive_NothingAdded_SupportedFunctionsSize(test *testing.T) {
	evaluator := NewEvaluator()
	functions := evaluator.GetFunctions()
	expectedSize := 0

	if len(functions) != expectedSize {
		test.Errorf("Wrong number of supported functions: %v instead of %v", len(functions), expectedSize)
	}
}

func TestEvaluator_Positive_NothingAdded_SupportedOperatorsSize(test *testing.T) {
	evaluator := NewEvaluator()
	operators := evaluator.GetOperators()
	expectedSize := 0

	if len(operators) != expectedSize {
		test.Errorf("Wrong number of supported operators: %v instead of %v", len(operators), expectedSize)
	}
}

func TestEvaluator_Positive_NothingAdded_EvaluationResult(test *testing.T) {
	evaluator := NewEvaluator()
	expression := "A B C"

	value, err := evaluator.Evaluate(expression)

	if err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if *value != expression {
		test.Errorf("Wrong result %v instead of %v", *value, expression)
	}
}

func TestEvaluator_Positive_AddFunction_SupportedFunctionsSize(test *testing.T) {
	evaluator := NewEvaluator()
	expectedSize := 1

	if err := evaluator.AddFunction(newDummyFunction()); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	functions := evaluator.GetFunctions()
	if len(functions) != expectedSize {
		test.Errorf("Wrong number of supported functions: %v instead of %v", len(functions), expectedSize)
	}
}

func TestEvaluator_Positive_AddOperator_SupportedOperatorsSize(test *testing.T) {
	evaluator := NewEvaluator()
	expectedSize := 1

	if err := evaluator.AddOperator(newDummyOperator()); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	operators := evaluator.GetOperators()
	if len(operators) != expectedSize {
		test.Errorf("Wrong number of supported operators: %v instead of %v", len(operators), expectedSize)
	}
}

func TestEvaluator_Positive_AddFunction_VerifyTokenizer(test *testing.T) {
	dummyFunction := newDummyFunction()

	tokenizerMock := &mockedTokenizer{}
	tokenizerMock.On("AddFunction", dummyFunction.Name()).Return(nil)
	defer tokenizerMock.AssertCalled(test, "AddFunction", dummyFunction.Name())

	evaluator := newMockedEvaluator(nil, tokenizerMock)
	assert.NoError(test, evaluator.AddFunction(dummyFunction))
}

func TestEvaluator_Positive_AddOperator_VerifyConverter(test *testing.T) {
	dummyOperator := newDummyOperator()

	converterMock := &mockedConverter{}
	converterMock.On("AddOperator", dummyOperator.Name(), dummyOperator.Priority()).Return(nil)
	defer converterMock.AssertCalled(test, "AddOperator", dummyOperator.Name(), dummyOperator.Priority())

	tokenizerMock := &mockedTokenizer{}
	tokenizerMock.On("AddOperator", dummyOperator.Name()).Return(nil)

	evaluator := newMockedEvaluator(converterMock, tokenizerMock)
	assert.NoError(test, evaluator.AddOperator(dummyOperator))
}

func TestEvaluator_Positive_AddOperator_VerifyTokenizer(test *testing.T) {
	dummyOperator := newDummyOperator()

	converterMock := &mockedConverter{}
	converterMock.On("AddOperator", dummyOperator.Name(), dummyOperator.Priority()).Return(nil)

	tokenizerMock := &mockedTokenizer{}
	tokenizerMock.On("AddOperator", dummyOperator.Name()).Return(nil)
	defer tokenizerMock.AssertCalled(test, "AddOperator", dummyOperator.Name())

	evaluator := newMockedEvaluator(converterMock, tokenizerMock)
	assert.NoError(test, evaluator.AddOperator(dummyOperator))
}

func TestEvaluator_Positive_AddStandardFunctions_SupportedFunctionsSize(test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddStandardFunctions(); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(evaluator.GetFunctions()) == 0 {
		test.Errorf("Wrong number of supported functions: 0")
	}
}

func TestEvaluator_Positive_AddStandardOperators_SupportedOperatorsSize(test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddStandardOperators(); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if len(evaluator.GetOperators()) == 0 {
		test.Errorf("Wrong number of supported operators: 0")
	}
}

func TestEvaluator_Positive_Evaluate_VerifyConverter(test *testing.T) {
	dummyOperator := newDummyOperator()
	expression := "A " + dummyOperator.Name() + " B"
	infixTokens := []token{{atomType, "A"}, {operatorType, dummyOperator.Name()}, {atomType, "B"}}
	postfixTokens := []token{{atomType, "A"}, {atomType, "B"}, {operatorType, dummyOperator.Name()}}

	converterMock := &mockedConverter{}
	converterMock.On("AddOperator", dummyOperator.Name(), dummyOperator.Priority()).Return(nil)
	converterMock.On("Convert", infixTokens).Return(postfixTokens, nil)
	defer converterMock.AssertCalled(test, "Convert", infixTokens)

	tokenizer := newInfixTokenizer()

	evaluator := newMockedEvaluator(converterMock, tokenizer)
	assert.NoError(test, evaluator.AddOperator(dummyOperator))

	_, err := evaluator.Evaluate(expression)
	assert.NoError(test, err)
}

func TestEvaluator_Positive_Evaluate_VerifyTokenizer(test *testing.T) {
	dummyOperator := newDummyOperator()
	expression := "A " + dummyOperator.Name() + " B"
	infixTokens := []token{{atomType, "A"}, {operatorType, dummyOperator.Name()}, {atomType, "B"}}

	converter := newInfixConverter()

	tokenizerMock := &mockedTokenizer{}
	tokenizerMock.On("AddOperator", dummyOperator.Name()).Return(nil)
	tokenizerMock.On("Tokenize", expression).Return(infixTokens, nil)
	defer tokenizerMock.AssertCalled(test, "Tokenize", expression)

	evaluator := newMockedEvaluator(converter, tokenizerMock)
	assert.NoError(test, evaluator.AddOperator(dummyOperator))

	_, err := evaluator.Evaluate(expression)
	assert.NoError(test, err)
}

func TestEvaluator_Positive_OneOperator_EvaluationResult(test *testing.T) {
	expression := "A + B"
	expectedValue := "AB"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_TwoOperator_EvaluationResult(test *testing.T) {
	expression := "A + B + C"
	expectedValue := "ABC"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_OneFunction_EvaluationResult(test *testing.T) {
	expression := "IF(A < B, 1, 2)"
	expectedValue := "1"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_NestedFunction_EvaluationResult(test *testing.T) {
	expression := "IF(A > B, 1, REPLACE(ABC, A, E))"
	expectedValue := "EBC"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_TwoFunctionsOneOperator_EvaluationResult(test *testing.T) {
	expression := "IF(A < B, 1, 2) + IF(A > B, 3, 4)"
	expectedValue := "14"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_RedundantBrackets_EvaluationResult(test *testing.T) {
	expression := "(((A + B)))"
	expectedValue := "AB"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_CompareWithNull_EvaluationResult(test *testing.T) {
	expression := "IF(A == NULL, Wrong, Correct)"
	expectedValue := "Correct"
	checkEvaluatorResult(expression, expectedValue, test)
}

func TestEvaluator_Positive_NullAsResult_EvaluationResult(test *testing.T) {
	expression := "IF(A == B, Wrong, NULL)"

	evaluator := NewEvaluator()

	if err := evaluator.AddStandardFunctions(); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := evaluator.AddStandardOperators(); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	value, err := evaluator.Evaluate(expression)
	if err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if value != nil {
		test.Errorf("Wrong value %v instead of %v", value, nil)
	}
}

func TestEvaluator_Negative_AddNullFunction(test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddFunction(nil); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Evaluator: added function is empty" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestEvaluator_Negative_AddNullOperator(test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddOperator(nil); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Evaluator: added operator is empty" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestEvaluator_Negative_TwoFunctionsWithTheSameName(test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddFunction(newDummyFunction()); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := evaluator.AddFunction(newDummyFunction()); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Evaluator: function "+newDummyFunction().Name()+" is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestEvaluator_Negative_TwoOperatorsWithTheSameName(test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddOperator(newDummyOperator()); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if err := evaluator.AddOperator(newDummyOperator()); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != "Evaluator: operator "+newDummyOperator().Name()+" is already added" {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}

func TestEvaluator_Negative_UnpairedBracket(test *testing.T) {
	expression := "A + (B + C"
	expectedError := "Converter: unpaired bracket"
	checkEvaluatorError(expression, expectedError, test)
}

func TestEvaluator_Negative_FewArguments(test *testing.T) {
	expression := "A + "
	expectedError := "BaseOperator: not enough arguments for operator +"
	checkEvaluatorError(expression, expectedError, test)
}

func TestEvaluator_Negative_FewOperators(test *testing.T) {
	expression := "A + B C"
	expectedError := "Evaluator: invalid expression"
	checkEvaluatorError(expression, expectedError, test)
}

func TestEvaluator_Negative_Evaluate_UnexpectedTokenType(test *testing.T) {
	expression := "A + B"
	wrongTokens := []token{{atomType, "A"}, {atomType, "B"}, {leftBracketType, "("}}

	converterMock := &mockedConverter{}
	converterMock.On("AddOperator", mock.Anything, mock.Anything).Return(nil)
	converterMock.On("Convert", mock.Anything).Return(wrongTokens, nil)

	tokenizer := newInfixTokenizer()

	evaluator := newMockedEvaluator(converterMock, tokenizer)
	assert.NoError(test, evaluator.AddStandardFunctions())
	assert.NoError(test, evaluator.AddStandardOperators())

	_, err := evaluator.Evaluate(expression)
	assert.Error(test, err)
	assert.Equal(test, err.Error(), fmt.Sprintf("Evaluator: unexpected token type %v", leftBracketType))
}

func TestEvaluator_Negative_Evaluate_UnsupportedFunction(test *testing.T) {
	expression := "A + B"
	wrongTokens := []token{{atomType, "A"}, {atomType, "B"}, {functionType, "FUNC"}}

	converterMock := &mockedConverter{}
	converterMock.On("AddOperator", mock.Anything, mock.Anything).Return(nil)
	converterMock.On("Convert", mock.Anything).Return(wrongTokens, nil)

	tokenizer := newInfixTokenizer()

	evaluator := newMockedEvaluator(converterMock, tokenizer)
	assert.NoError(test, evaluator.AddStandardFunctions())
	assert.NoError(test, evaluator.AddStandardOperators())

	_, err := evaluator.Evaluate(expression)
	assert.Error(test, err)
	assert.Equal(test, err.Error(), fmt.Sprintf("Evaluator: unsupported function FUNC"))
}

func TestEvaluator_Negative_Evaluate_UnsupportedOperator(test *testing.T) {
	expression := "A + B"
	wrongTokens := []token{{atomType, "A"}, {atomType, "B"}, {operatorType, "<>"}}

	converterMock := &mockedConverter{}
	converterMock.On("AddOperator", mock.Anything, mock.Anything).Return(nil)
	converterMock.On("Convert", mock.Anything).Return(wrongTokens, nil)

	tokenizer := newInfixTokenizer()

	evaluator := newMockedEvaluator(converterMock, tokenizer)
	assert.NoError(test, evaluator.AddStandardFunctions())
	assert.NoError(test, evaluator.AddStandardOperators())

	_, err := evaluator.Evaluate(expression)
	assert.Error(test, err)
	assert.Equal(test, err.Error(), fmt.Sprintf("Evaluator: unsupported operator <>"))
}

type dummyFunction struct {
	functions.BaseFunction
}

func newDummyFunction() *dummyFunction {
	result := &dummyFunction{functions.MakeBaseFunction(nil, "Function", 2)}
	result.SetDerived(result)
	return result
}

func (f *dummyFunction) CheckArguments(arguments []interface{}) bool {
	return true
}

func (f *dummyFunction) Result(arguments []interface{}) interface{} {
	return "FunctionResult"
}

type dummyOperator struct {
	operators.BaseOperator
}

func newDummyOperator() *dummyOperator {
	result := &dummyOperator{operators.MakeBaseOperator(nil, "Operator", 1, 2)}
	result.SetDerived(result)
	return result
}

func (f *dummyOperator) CheckArguments(arguments []interface{}) bool {
	return true
}

func (f *dummyOperator) Result(arguments []interface{}) interface{} {
	return "OperatorResult"
}

type mockedConverter struct {
	mock.Mock
}

func (m *mockedConverter) AddOperator(name string, priority int) error {
	args := m.Called(name, priority)
	return args.Error(0)
}

func (m *mockedConverter) Convert(tokens []token) ([]token, error) {
	args := m.Called(tokens)
	return args.Get(0).([]token), args.Error(1)
}

type mockedTokenizer struct {
	mock.Mock
}

func (m *mockedTokenizer) AddFunction(function string) error {
	args := m.Called(function)
	return args.Error(0)
}

func (m *mockedTokenizer) AddOperator(operator string) error {
	args := m.Called(operator)
	return args.Error(0)
}

func (m *mockedTokenizer) Tokenize(expression string) ([]token, error) {
	args := m.Called(expression)
	return args.Get(0).([]token), args.Error(1)
}
