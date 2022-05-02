package main

import (
	"strconv"
	"strings"
)

const OperatorSymbols = ",.-+*/^"
const allowedChars = "abcdefghijklmnopqrstuvwxyz"

var index int = 0

//"<", "(", ")", ".", "=", "clr", "1", "2", "3", "+", "-", "fbr", "4", "5", "6", "*", "/", "√", "7", "8", "9", "0", "^", "³√"
func HandleInput(Input string) {
	LogMessage("Tried adding: " + Input)
	switch {
	case Input == "⌫":
		MathExpression.CutSuffix()
		index = len(MathExpression.numbers) - 1
		updateText()
		return
	case Input == "clr":
		index = 0
		MathExpression.Clear()
		updateText()
		return
	case strings.ContainsAny(Input, AllowedSymbols): //if is symbol
		if len(MathExpression.numbers[index]) > 0 { //if there are any numbers in last expression
			MathExpression.Fix()
			index++
			MathExpression.operations = append(MathExpression.operations, Input)
			MathExpression.numbers = append(MathExpression.numbers, "") //enlarge size of numbers slice
		} else { //if there are none numbers in last expression
			MathExpression.operations[index] += Input
		}
	case isInt(Input): //numbers aren't function
		MathExpression.numbers[index] += Input
	case Input == "." || Input == ",":
		if len(MathExpression.numbers[index]) == 0 {
			MathExpression.numbers[index] += "0"
		}
		MathExpression.numbers[index] += "."
	case Input == "=":
		Calculate()
	case Input == "(":
		if len(MathExpression.numbers[index]) != 0 { //if something is before the (
			HandleInput("*")
			HandleInput("(")
			return
		}
		MathExpression.numbers[index] += Input
	case Input == ")":
		if strings.Count(MathExpression.numbers[index], "(") > strings.Count(MathExpression.numbers[index], ")") {
			MathExpression.numbers[index] += Input
			MathExpression.CutEndBrackets()
		}
	// case Input == "√":
	// 	MathExpression += "sqrt("
	// 	LogMessage("Not yet implemented square root")
	// case Input == "³√":  //to bedzie cienzki dzien jak to bd robil
	// 	LogMessage("Not yet implemented cube root")
	default:
		return
	}
	updateText()
}

func containsChars(a string) bool {
	if strings.ContainsAny(a, allowedChars) {
		return true
	}
	return false
}

func isInt(a string) bool {
	if _, err := strconv.Atoi(a); err == nil {
		return true
	}
	return false
}

func PressButton(id int) {
	HandleInput(Operators[id])
}
