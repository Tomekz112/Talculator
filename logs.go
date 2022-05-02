package main

import (
	"errors"
	"os"
	"strconv"
	"time"
)

func LogError(errType, err string, shutdown bool) {
	currentTime := time.Now()
	l, oErr := os.OpenFile("Logs/"+currentTime.Format("2006.01.02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errors.Is(oErr, os.ErrNotExist) {
		os.WriteFile("Logs/"+currentTime.Format("2006.01.02")+".log", []byte(currentTime.Format("2006.01.02 15:04:05")+" "+errType+" has occurred: "+err+", shutdown: "+strconv.FormatBool(shutdown)+"\n"), 0644)
	} else {
		l.Write([]byte(currentTime.Format("2006.01.02 15:04:05") + " " + errType + " has occurred: " + err + ", shutdown: " + strconv.FormatBool(shutdown) + "\n"))
	}
	if oErr != nil || shutdown {
		os.Exit(3)
	}
	l.Close()
}

func LogMessage(message string) {
	currentTime := time.Now()
	l, oErr := os.OpenFile("Logs/"+currentTime.Format("2006.01.02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errors.Is(oErr, os.ErrNotExist) {
		os.WriteFile("Logs/"+currentTime.Format("2006.01.02")+".log", []byte(currentTime.Format("2006.01.02 15:04:05")+" "+message+"\n"), 0644)
	} else {
		l.Write([]byte(currentTime.Format("2006.01.02 15:04:05") + " " + message + "\n"))
	}
	l.Close()
}
