package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var str = `\`
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("Resources"+str+"rcedit.exe", path+str+"app.exe", "--set-icon", "Resources"+str+"icon.ico")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
	fmt.Println("Path:", path)
	fmt.Println("Press enter to exit")
	fmt.Scanln(&str)
}
