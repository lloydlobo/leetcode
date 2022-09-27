package helper

import (
	"fmt"
	"log"
)

func LogFile() {
	fmt.Println("\n ")
	log.SetFlags(3)
	log.Printf("| ~~~~main.go~~~~\n")
}

func ClearTerminalUnix() {
	fmt.Println("\033[2J")
}
