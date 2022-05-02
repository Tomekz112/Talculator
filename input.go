package main

import (
	"strconv"
	"strings"
)

const OperatorSymbols = ",.-+*/^()"
const AllowedFuncChars = "abcdefghijklmnopqrstuvwxyz"

var MathExpression Equation
var funcMathExpressions []Equation
var Functions []Function

//"<", "(", ")", ".", "=", "clr", "1", "2", "3", "+", "-", "fbr", "4", "5", "6", "*", "/", "√", "7", "8", "9", "0", "^", "³√"
func HandleInput(Input string) {
	LogMessage("Tried adding: " + Input)
	switch {
	case Input == "⌫":
		MathExpression.CutSuffix()
	case Input == "clr":
		MathExpression.Clear()
	case Input == "=":
		MathExpression.Calculate()
	case strings.ContainsAny(Input, OperatorSymbols) || IsInt(Input):
		MathExpression.Add(Input)
	case strings.ContainsAny(Input, AllowedFuncChars):
		MathExpression.AddFunc(Input)
	}
	updateText()
}

func IsInt(a string) bool {
	if _, err := strconv.Atoi(a); err == nil {
		return true
	}
	return false
}

func PressButton(id int) {
	HandleInput(Operators[id])
}
