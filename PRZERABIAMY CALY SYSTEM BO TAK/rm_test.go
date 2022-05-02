package main

import (
	"testing"
)

func removingDupe() {
	MathExpression.CutSuffix()
}

// test function
func TestRemovingText(t *testing.T) {
	var equation string = "222+2+4*2323*232+12312312*"
	var iteration int = 3
	MathExpression = ToEquation(equation)
	expectedString := equation
	for i := 0; i < iteration; i++ {
		MathExpression.CutSuffix()
		expectedString = expectedString[:len(expectedString)-1]
		t.Log("rot:", i+1, "exStr:", expectedString, "rlstring:", MathExpression.ToString())
		if MathExpression.ToString() != expectedString {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedString, MathExpression.ToString())
		} else {
			t.Logf("Removing characters test no (%x) passed", i+1)
		}
	}
}
