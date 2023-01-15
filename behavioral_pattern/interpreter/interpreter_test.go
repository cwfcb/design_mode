package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	expression := "8 3 2 4 - + *"
	expInterpreter := &ExpressionInterpreter{}
	result, err := expInterpreter.interpreter(expression)
	t.Error(err)
	t.Log(result)
}
