package main

import (
	"fmt"
	"io/ioutil"
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
		file, err := os.Open(Args[2])
		defer file.Close()
		if err != nil {
			fmt.Printf(Red + "[ERROR]File Not Found" + Reset + "\n")
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf(Red + "[ERROR]Fail To Read File" + Reset + "\n")
		}
		if Args[1] == "H" {
			fmt.Printf("%x\n", data)
		} else if Args[1] == "S" {
			fmt.Printf("%s\n", data)
		}
		//fmt.Println("----------------Results----------------")
		fmt.Println("Number of bytes read:", len(data))
	}
}
