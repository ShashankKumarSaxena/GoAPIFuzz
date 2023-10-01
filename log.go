package main

import (
	"fmt"
	"strings"
)

func PrintLog(logType string, message string) {
	fmt.Println("[" + strings.ToUpper(logType) + "]" + " " + message)
}
