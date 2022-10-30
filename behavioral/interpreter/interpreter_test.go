package interpreter

import "testing"

func TestAddExpression_Interpret(t *testing.T) {
	if DoExpression("add", Context{1}, Context{2}) != 3 {
		t.Error("1+2!=3")
	}
	if DoExpression("sub", Context{2}, Context{1}) != 1 {
		t.Error("2-1!=1")
	}
}
