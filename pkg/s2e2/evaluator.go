package s2e2

import (
	"fmt"
	"github.com/mzinin/s2e2.go/pkg/s2e2/functions"
	"github.com/mzinin/s2e2.go/pkg/s2e2/operators"
)

const (
	// Null value in an input expression.
	nullValue = "NULL"

	// Expected stack size after processing all tokens.
	finalStackSize = 1
)

// Evaluator evaluates string value of an expression.
type Evaluator struct {
	cnvtr     converter                     // Converter of infix token sequence into postfix one.
	tknzr     tokenizer                     // Tokenizer of expression into list of tokens.
	functions map[string]functions.Function // Set of all supported functions.
	operators map[string]operators.Operator // Set of all supported operators.
	stack     []interface{}                 // Stack of intermediate values.
}

// NewEvaluator creates an instance of Evaluator.
func NewEvaluator() *Evaluator {
	result := &Evaluator{}
	result.cnvtr = newInfixConverter()
	result.tknzr = newInfixTokenizer()
	result.functions = make(map[string]functions.Function)
	result.operators = make(map[string]operators.Operator)
	result.stack = make([]interface{}, 0)
	return result
}

// newMockedEvaluator creates an instance of Evaluator with external converter and tokenizer.
func newMockedEvaluator(cnvtr converter, tknzr tokenizer) *Evaluator {
	result := &Evaluator{}
	result.cnvtr = cnvtr
	result.tknzr = tknzr
	result.functions = make(map[string]functions.Function)
	result.operators = make(map[string]operators.Operator)
	result.stack = make([]interface{}, 0)
	return result
}

// AddFunction adds function to set of supported functions.
// Returns error if function is nil or if function or operator with the same name is already added.
func (e *Evaluator) AddFunction(function functions.Function) error {
	if function == nil {
		return fmt.Errorf("Evaluator: added function is empty")
	}
	if err := e.checkUniqueness(function.Name()); err != nil {
		return err
	}

	if err := e.tknzr.AddFunction(function.Name()); err != nil {
		return err
	}
	e.functions[function.Name()] = function
	return nil
}

// AddOperator add operator to set of supported operators.
// Returns error if operator is nil or if function or operator with the same name is already added.
func (e *Evaluator) AddOperator(operator operators.Operator) error {
	if operator == nil {
		return fmt.Errorf("Evaluator: added operator is empty")
	}
	if err := e.checkUniqueness(operator.Name()); err != nil {
		return err
	}

	if err := e.cnvtr.AddOperator(operator.Name(), operator.Priority()); err != nil {
		return err
	}
	if err := e.tknzr.AddOperator(operator.Name()); err != nil {
		return err
	}
	e.operators[operator.Name()] = operator
	return nil
}

// AddStandardFunctions adds all standard functions to set of supported functions.
// Returns error if there is a collision between functions names.
func (e *Evaluator) AddStandardFunctions() error {
	if err := e.AddFunction(functions.NewFunctionAddDays()); err != nil {
		return err
	}
	if err := e.AddFunction(functions.NewFunctionFormatDate()); err != nil {
		return err
	}
	if err := e.AddFunction(functions.NewFunctionIf()); err != nil {
		return err
	}
	if err := e.AddFunction(functions.NewFunctionNow()); err != nil {
		return err
	}
	if err := e.AddFunction(functions.NewFunctionReplace()); err != nil {
		return err
	}
	return nil
}

// AddStandardOperators adds all standard operators to set of supported operators.
// Returns error if there is a collision between operators names.
func (e *Evaluator) AddStandardOperators() error {
	if err := e.AddOperator(operators.NewOperatorAnd()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorEqual()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorGreaterOrEqual()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorGreater()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorLessOrEqual()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorLess()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorNotEqual()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorNot()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorOr()); err != nil {
		return err
	}
	if err := e.AddOperator(operators.NewOperatorPlus()); err != nil {
		return err
	}

	return nil
}

// GetFunctions gets collection of all supported functions.
func (e *Evaluator) GetFunctions() []functions.Function {
	result := make([]functions.Function, 0, len(e.functions))
	for _, function := range e.functions {
		result = append(result, function)
	}
	return result
}

// GetOperators gets collection of all supported operators.
func (e *Evaluator) GetOperators() []operators.Operator {
	result := make([]operators.Operator, 0, len(e.operators))
	for _, operator := range e.operators {
		result = append(result, operator)
	}
	return result
}

// Evaluate get the value of the input expression.
// Returns error in case of an invalid expression.
func (e *Evaluator) Evaluate(expression string) (*string, error) {
	infixExpression, err := e.tknzr.Tokenize(expression)
	if err != nil {
		return nil, err
	}

	// a bit of syntax sugar: if expression contains only atoms
	// consider it as just a string literal
	if e.onlyAtoms(infixExpression) {
		return &expression, nil
	}

	postfixExpression, err := e.cnvtr.Convert(infixExpression)
	if err != nil {
		return nil, err
	}

	return e.evaluateExpression(postfixExpression)
}

// checkUniqueness checks is function's or operator's name is unique.
// Returns error if the name is not unique.
func (e *Evaluator) checkUniqueness(entityName string) error {
	if _, ok := e.functions[entityName]; ok {
		return fmt.Errorf("Evaluator: function %v is already added", entityName)
	}
	if _, ok := e.operators[entityName]; ok {
		return fmt.Errorf("Evaluator: operator %v is already added", entityName)
	}
	return nil
}

// onlyAtoms checks is all tokens in the slice are atoms
func (e *Evaluator) onlyAtoms(tokens []token) bool {
	result := true
	for _, t := range tokens {
		if t.Type != atomType {
			result = false
			break
		}
	}
	return result
}

// evaluateExpression gets value of the postfix sequence of tokens.
// Returns error expression is invalid.
func (e *Evaluator) evaluateExpression(postfixExpression []token) (*string, error) {
	e.stack = make([]interface{}, 0)

	for _, tkn := range postfixExpression {
		switch tkn.Type {
		case atomType:
			e.processAtom(tkn)

		case functionType:
			if err := e.processFunction(tkn); err != nil {
				return nil, err
			}

		case operatorType:
			if err := e.processOperator(tkn); err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("Evaluator: unexpected token type %v", tkn.Type)
		}
	}

	result, err := e.getResultValueFromStack()
	if err != nil {
		return nil, err
	}

	e.stack = nil
	return result, nil
}

// processAtom processes ATOM token.
func (e *Evaluator) processAtom(tkn token) {
	if tkn.Value == nullValue {
		e.stack = append(e.stack, nil)
	} else {
		e.stack = append(e.stack, tkn.Value)
	}
}

// processFunction processes FUNCTION token.
// Returns error is case of unknown function.
func (e *Evaluator) processFunction(tkn token) error {
	function, ok := e.functions[tkn.Value]

	if !ok {
		return fmt.Errorf("Evaluator: unsupported function %v", tkn.Value)
	}
	return function.Invoke(&e.stack)
}

// processOperator processes OPERATOR token.
// Returns error is case of unknown operator.
func (e *Evaluator) processOperator(tkn token) error {
	operator, ok := e.operators[tkn.Value]

	if !ok {
		return fmt.Errorf("Evaluator: unsupported operator %v", tkn.Value)
	}
	return operator.Invoke(&e.stack)
}

// getResultValueFromStack get result value from the stack of intermediate values.
func (e *Evaluator) getResultValueFromStack() (*string, error) {
	if len(e.stack) != finalStackSize {
		return nil, fmt.Errorf("Evaluator: invalid expression")
	}

	if e.stack[0] == nil {
		return nil, nil
	}

	result, ok := e.stack[0].(string)
	if !ok {
		return nil, fmt.Errorf("Evaluator: expression value is not a string")
	}
	return &result, nil
}
