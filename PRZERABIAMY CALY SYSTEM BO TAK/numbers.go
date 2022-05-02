package main

import (
	"strings"
)

var AllowedSymbols string = "+-/*^³√"

type Equation struct {
	operations []string
	numbers    []string
	brackets   string
}

var EmptyEquation = Equation{operations: []string{""}, numbers: []string{""}, brackets: ""}

func (eq *Equation) Clear() {
	*eq = Equation{operations: []string{""}, numbers: []string{""}, brackets: ""}
}

func (eq *Equation) Normalize() {
	for i := 0; i < len(eq.numbers); i++ {
		eq.operations[i] = strings.Replace(eq.operations[i], "^", "**", -1)
	}
}

func (eq *Equation) GetLastCharIndex() int {
	length := len(eq.numbers)
	return len(eq.numbers[length-1]) - 1
}

func (eq *Equation) GetLastOperationIndex() int {
	length := len(eq.operations)
	return len(eq.operations[length-1]) - 1
}

func (eq *Equation) LastChar() string {
	index := len(eq.numbers) - 1
	if len(eq.numbers[index]) == 0 {
		return ""
	}
	return string(eq.numbers[index][len(eq.numbers[index])-1])
}

// func (eq *Equation) LastOperation() string {
// 	index := len(eq.operations) - 1
//if len(eq.operations[index]) == 0 {
//return ""
//}
// 	return string(eq.operations[index][len(eq.operations[index])-1])
// }

//Fix fixes the last equation by:
//adding closing brackets,
//removing . as last character in number
//seting 0 as first character if . was first before
func (eq *Equation) Fix() {
	eq.Normalize()
	if eq.LastChar() == "." {
		eq.CutSuffix()
	}
	//Remove empty strings from slices
	oldEq := *eq
	*eq = Equation{operations: []string{""}, numbers: []string{}, brackets: ""}
	for _, str := range oldEq.operations {
		if str != "" {
			eq.operations = append(eq.operations, str)
		}
	}
	for _, str := range oldEq.numbers {
		if str != "" {
			eq.numbers = append(eq.numbers, str)
		}
	}
	if len(eq.numbers) == 0 {
		eq.numbers = append(eq.numbers, "")
	}
	//End of the empty strings check
	//Check if there are any missing brackets
	strNumbers := strings.Join(eq.numbers, "")
	neededbrackets := strings.Count(strNumbers, "(") - strings.Count(strNumbers+eq.brackets, ")")
	for i := 0; i != neededbrackets; i++ {
		eq.brackets += ")"
	}
	//End of the bracket check
}

func (eq *Equation) CutEndBrackets() {
	eq.brackets = strings.TrimSuffix(eq.brackets, ")")
}

func NormalizeString(str string) string {
	str = strings.Replace(str, ",", ".", -1)
	str = strings.Replace(str, "^", "**", -1)
	return str
}

func (eq *Equation) CutSuffix() {
	nLen := len(eq.numbers)
	opLen := len(eq.operations)
	if nLen == 0 || eq.numbers[0] == "" { //if its empty
		return
	}
	if opLen > nLen { //remove operation
		eq.operations[opLen-1] = eq.operations[opLen-1][:eq.GetLastOperationIndex()]
	} else { //remove number or bracket
		if eq.LastChar() == "(" {
			eq.CutEndBrackets()
		}
		eq.numbers[nLen-1] = eq.numbers[nLen-1][:eq.GetLastCharIndex()]
	}
	eq.Fix()
}

//BACKUP func (eq *Equation) CutSuffix() {
// 	length := len(eq.numbers)
// 	if eq.numbers[0] == "" {
// 		return
// 	}
// 	if len(eq.operations) == length && eq.operations[length-1] != "" {
// 		fmt.Println("a")
// 		eq.operations = eq.operations[:length-1]
// 	} else {
// 		if string(eq.numbers[length-1][eq.GetLastCharIndex()]) == "(" {
// 			eq.CutEndBrackets()
// 		} else if string(eq.numbers[length-1][eq.GetLastCharIndex()]) == ")" {
// 			eq.brackets += ")"
// 		}
// 		eq.numbers[length-1] = eq.numbers[length-1][:eq.GetLastCharIndex()]
// 	}
//}

func (eq *Equation) ToString() string {
	eq.Fix()
	var output string
	for i := 0; i < len(eq.numbers); i++ {
		output += eq.operations[i] + eq.numbers[i]
	}
	if len(eq.operations) > len(eq.numbers) {
		output += eq.operations[len(eq.operations)-1]
	}
	return output + eq.brackets
}

//ToEquation converts string to equation and return it,
//The Equation should be Normalized before
func ToEquation(str string) Equation {
	var eq = Equation{}
	eq.operations = append(eq.operations, "")
	var adchars string //bad initialization of negative numbers
	for i := strings.IndexAny(str, AllowedSymbols); i != -1; i = strings.IndexAny(str, AllowedSymbols) {
		if i != 0 {
			eq.operations = append(eq.operations, string(str[i])) //sets number as everything before mathematic symbol
			eq.numbers = append(eq.numbers, adchars+str[:i])
			adchars = ""
		} else if string(str[i]) != "-" { //it's not -
			eq.operations[len(eq.operations)-1] += string(str[i]) //adds the symbol to previous one
		} else {
			adchars = "-" //bad initialization of negative numbers
		}
		str = str[i+1:] //cuts added elements from the string
	}
	eq.numbers = append(eq.numbers, adchars+str)
	eq.Fix()
	return eq
}
