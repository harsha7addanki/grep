package main

import (
	"fmt"
	"os"
	"runtime"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func main() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
	Args := os.Args[1:]
	/*
		fmt.Printf("/---------------\\          /---------------          /---------------\n")
		fmt.Printf("|                |          |              |          |               \n")
		fmt.Printf("|                           |              /          |               \n")
		fmt.Printf("|           ---\\           |-------------/           |---------------\n")
		fmt.Printf("\\             |            |\\                       |               \n")
		fmt.Printf(" \\            |            | \\                      |               \n")
		fmt.Printf("  \\-----------/            |  \\                     |---------------\n")
	*/
	if Args[0] == "READ" {
		file, err := os.Open(Args[1])
		defer file.Close()
		if err != nil {
			fmt.Printf(Red + "[ERROR]File Not Found" + Reset + "\n")
		}
	}
}
