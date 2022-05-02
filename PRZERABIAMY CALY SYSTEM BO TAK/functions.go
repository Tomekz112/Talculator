// package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// type Function struct {
// 	name, Desc, info string
// 	Input            []int
// 	active           bool
// }

// func main() {
// 	fmt.Println("Starting...")
// 	var function Function = Function{
// 		name:   "test",
// 		active: false,
// 	}
// 	function.Load()
// }

// func (f *Function) ToString() string {
// 	var s string = "("
// 	for i := 0; i < len(f.Input); i++ {
// 		s += strconv.Itoa(f.Input[i]) + ","
// 	}
// 	return f.name + s + ")"
// }

// func (f *Function) Load() {
// 	properties, err := os.ReadFile("Functions/" + f.name + ".json")
// 	if errors.Is(err, os.ErrNotExist) {
// 		f.info = "Didn't found any function with given name! Please check the seplling"
// 		return
// 	} else if err != nil {
// 		return
// 	}
// 	json.Unmarshal(properties, &f)
// 	f.info = "Function found, please enter: " + strconv.Itoa(len(f.Input)) + " values"
// 	fmt.Println(f.info)
// 	fmt.Println(f.Desc)
// }

// func (f *Function) SolveFunction() string {
// 	fmt.Println("Not implemented yet")
// 	return ""
// }
package main

import (
	"MyPkgs/toxel"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Function struct {
	Name        string
	Description string
	Numbers     int
}

var TextActive = false

//Returns errorcode
//0 - function passed correctly
//1 - couldn't find function in directory
//2 - other error
func SearchFunction(name string) (Function, int) {
	var function Function
	fmt.Println("Functions/" + name + ".json")
	properties, err := os.ReadFile("Functions/" + name + ".json")
	if errors.Is(err, os.ErrNotExist) {
		return function, 1
	} else if err != nil {
		return function, 2
	}
	json.Unmarshal(properties, &function)
	return function, 0
}

var Message *toxel.ToxelText

func HandleNewFunc() {
	function, errCode := SearchFunction(MathExpression.numbers[len(MathExpression.numbers)-1])
	switch errCode {
	case 0:
		createTextBox("Found " + function.Name + " function!\n Please enter " + strconv.Itoa(function.Numbers) +
			" values\n" + function.Name + " - " + function.Description)

	case 1:
		createTextBox("Didn't found any function with given name! Please check the seplling")
	case 2:
		createTextBox("Error!")
	}
}

func createTextBox(text string) {
	TextActive = true
	Message.Text = text
}
