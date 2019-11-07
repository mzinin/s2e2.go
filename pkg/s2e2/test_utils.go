package s2e2

import "testing"

func compareTokens(actual, expected []token, test *testing.T) {
	if len(actual) != len(expected) {
		test.Errorf("Wrong output size %v instead of %v", len(actual), len(expected))
		return
	}
	for i, value := range actual {
		if value != expected[i] {
			test.Errorf("Wrong output token %v instead of %v", value, expected[i])
			return
		}
	}
}

func checkEvaluatorResult(expression, expectedValue string, test *testing.T) {
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

	if *value != expectedValue {
		test.Errorf("Wrong value %v instead of %v", *value, expectedValue)
	}
}

func checkEvaluatorError(expression, expectedError string, test *testing.T) {
	evaluator := NewEvaluator()

	if err := evaluator.AddStandardFunctions(); err != nil {
		test.Errorf("Unexpected error %v", err)
	}
	if err := evaluator.AddStandardOperators(); err != nil {
		test.Errorf("Unexpected error %v", err)
	}

	if _, err := evaluator.Evaluate(expression); err == nil {
		test.Errorf("No expected error")
	} else if err.Error() != expectedError {
		test.Errorf("Unexpected error text: %q", err.Error())
	}
}
