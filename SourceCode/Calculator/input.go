package main

import (
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/faiface/pixel/pixelgl"
)

const OperatorSymbols = ",.-+*/^"

//"<", "(", ")", ".", "=", "clr", "1", "2", "3", "+", "-", "fbr", "4", "5", "6", "*", "/", "√", "7", "8", "9", "0", "^", "³√"
func HandleInput(Input string) {
	var tempSuffix string = "0"
	switch {
	case Input == "=":
		Calculate(MathExpression)
		return
	case Input == "(":
		if len(MathExpression) != 0 && isInt(string(MathExpression[len(MathExpression)-1])) {
			Input = "*("
		}
	case Input == ")":
		tempSuffix = ""
	case strings.ContainsAny(Input, OperatorSymbols): //check if Input is any operator
	case window.JustPressed(pixelgl.KeyBackspace) || Input == "<":
		MathExpression = cutSuffix(MathExpression) //remove last character from MathExpression
		updateText()
		return
	case isInt(Input):
	case Input == "clr":
		MathExpression = ""
		updateText()
		return
	case Input == "fbr":
		LogMessage("Reverted to fabricate settings")
	case Input == "√":
		MathExpression += "sqrt("
		LogMessage("Not yet implemented square root")
	case Input == "³√":  //to bedzie cienzki dzien jak to bd robil
		LogMessage("Not yet implemented cube root")
	default:
		return
	}
	if isCorrectExpression(MathExpression + Input + tempSuffix) { // +"0" is temp number after math operation
		MathExpression += Input
	}
	updateText()
}

func isCorrectExpression(mExpression string) bool {
	mExpression = StandardizeExpression(mExpression)
	if strings.Contains(mExpression, "(") && !strings.Contains(mExpression, ")") {
		mExpression += ")"
	}
	expression, err := govaluate.NewEvaluableExpression(mExpression)
	if err != nil {
		return false
	}
	_, err = expression.Evaluate(nil)
	return err == nil
}

func isInt(a string) bool {
	if _, err := strconv.Atoi(a); err == nil {
		return true
	}
	return false
}

func cutSuffix(s string) string {
	length := len(s)
	if length > 0 {
		s = s[:length-1]
	}
	return s
}

func PressButton(id int) {
	HandleInput(Operators[id])
}
