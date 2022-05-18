package main

import (
	"strconv"
	"strings"
)

const OperatorSymbols = ",.-+*/^()"
const AllowedFuncChars = "abcdefghijklmnopqrstuvwxyz"

var MathExpression Equation
var CurrentExpression *Equation
var Functions []Function
var CurrFunc int
var Depth int = 0

//"<", "(", ")", ".", "=", "clr", "1", "2", "3", "+", "-", "fbr", "4", "5", "6", "*", "/", "√", "7", "8", "9", "0", "^", "³√"
func HandleInput(Input string) {
	if Input == "" {
		return
	}
	LogMessage("Tried adding: " + Input)
	switch {
	case Input == "⌫":
		CurrentExpression.CutSuffix()
	case Input == "clr":
		MathExpression.Clear()
	case Input == "=":
		CurrentExpression.Calculate()
	case strings.ContainsAny(Input, OperatorSymbols) || IsInt(Input):
		CurrentExpression.Add(Input)
	case strings.ContainsAny(Input, AllowedFuncChars):
		CurrentExpression.AddFunc(Input)
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
