package main

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
)

var AllowedSymbols string = "+-/*^³√"
var fraction bool = false

type Equation struct {
	body     string
	brackets string
}

var EmptyEquation = Equation{body: "", brackets: ""}

func NormalizeString(str string) string {
	str = strings.Replace(str, ",", ".", -1)
	str = strings.Replace(str, "^", "**", -1)
	return str
}

func (eq *Equation) Body() string {
	return eq.body
}

func (eq *Equation) Brackets() string {
	return eq.brackets
}

func (eq *Equation) lastCharInt() bool {
	return IsInt(eq.lastChar()) || eq.lastChar() == ")"
}

func (eq *Equation) lastChar() string {
	if len(eq.body) == 0 {
		return ""
	}
	return string(eq.body[len(eq.body)-1])
}

func (eq *Equation) Clear() {
	eq = &EmptyEquation
}

func (eq *Equation) CutSuffix() {
	lastIndex := len(eq.body) - 1
	if lastIndex >= 0 {
		switch eq.body[lastIndex] {
		case '(':
			eq.brackets = strings.TrimSuffix(eq.brackets, ")")
		case ')':
			eq.brackets += ")"
		}
		eq.body = eq.body[:lastIndex]
	}
}

func (eq *Equation) Add(s string) {
	fmt.Println("Tried adding:", s)
	switch s {
	case "(":
		eq.brackets += ")"
		if eq.lastCharInt() {
			eq.body += "*"
		} else if strings.ContainsAny(eq.lastChar(), AllowedFuncChars) {
			funcName := eq.body[len(strings.TrimRight(eq.body, AllowedFuncChars)):]
			f, _ := SearchFunction(funcName, *eq)
			Functions = append(Functions, f)
		}
	case ")":
		if strings.Count(eq.body, "(") <= strings.Count(eq.body, ")") {
			return
		}
		eq.brackets = strings.TrimSuffix(eq.brackets, ")")
	case ".":
		fallthrough
	case ",":
		if fraction { //fraction - ułamek
			return
		}
		s = "."
		if !eq.lastCharInt() {
			eq.Add("0")
		}
		fraction = true
	default:
		if !IsInt(s) {
			if !eq.lastCharInt() {
				if s != "-" {
					return
				}
				eq.Add("(")
			}
			fraction = false
		}
	}
	eq.body += s
}

func (eq *Equation) AddFunc(s string) {
	if eq.lastCharInt() {
		eq.Add("*")
	}
	eq.body += s
}

//May be unsafe
func (eq *Equation) Set(s string) {
	eq.Clear()
	eq.body = s
}

func (eq *Equation) ToString() string {
	return eq.body + eq.brackets
}

func (eq *Equation) Calculate() {
	if !eq.lastCharInt() {
		eq.CutSuffix()
	}
	mExpression := eq.ToString()
	LogMessage("Calculating: " + mExpression)
	expression, err := govaluate.NewEvaluableExpression(mExpression)
	if err != nil {
		return
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return
	}
	strMathExpression := fmt.Sprintf("%v", result)
	LogMessage("Result: " + strMathExpression)
	strMathExpression = strings.Replace(strMathExpression, "Inf", "∞", -1)
	strMathExpression = NormalizeString(strMathExpression)
	eq.Set(strMathExpression)
	updateText()
}
