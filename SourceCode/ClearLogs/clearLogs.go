package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments, please use --help for info")
		return
	}
	if args[0] == "--help" {
		fmt.Println("This program removes ToxelCalc logs from current directory\nTo remove all files - run with all argument\nTo remove logs older than given days - run with number of days as argument")
	} else if args[0] == "all" {
		fmt.Println("Removing all logs")
		removeLogs(0)
	} else if isNumber(args[0]) {
		fmt.Println("Removing logs older than:", args[0], "days")
		i, _ := strconv.Atoi(args[0])
		removeLogs(i)
	} else {
		fmt.Println("Unknown arguments, please use --help")
	}
}

func removeLogs(days int) {
	files, err := ioutil.ReadDir("./Logs")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".log") {
			continue
		}
		if time.Since(file.ModTime()) < time.Hour*time.Duration(24*days) {
			return
		}
		e := os.Remove("Logs/" + file.Name())
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println("Removed:", file.Name())
	}
	fmt.Println("succesfully removed logs!")
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
