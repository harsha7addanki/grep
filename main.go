package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	fmt.Printf("/---------------\\          /---------------          /---------------         /---------\\\n")
	fmt.Printf("|               |          |              |          |                        |         |\n")
	fmt.Printf("|                          |              /          |                        |         |\n")
	fmt.Printf("|          ---\\            |_____________/           |---------------         |________/\n")
	fmt.Printf("\\             |            |\\                        |                        |\n")
	fmt.Printf(" \\            |            | \\                       |                        |\n")
	fmt.Printf("  \\-----------/            |  \\                      |---------------         |\n")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
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
		fmt.Println("----------------Results----------------")
		fmt.Println("Number Of Bytes Read:", len(data))
	} else if Args[0] == "FIND" {
		file, err := os.Open(Args[1])
		defer file.Close()
		if err != nil {
			fmt.Printf(Red + "[ERROR]File Not Found" + Reset + "\n")
		}
		if err != nil {
			fmt.Printf(Red + "[ERROR]Fail To Read File" + Reset + "\n")
		}
		linetip := bufio.NewScanner(file)
		var lines int = 0
		re := regexp.MustCompile(Args[2])
		finds := [][]int{}
		findlines := []string{}
		for {
			if linetip.Scan() == false {
				err = linetip.Err()
				if err == nil {
					break
				} else {
					fmt.Printf(Red + "[ERROR]Fail To Scan File" + Reset + "\n")
				}
			}
			if len(re.FindAllIndex([]byte(linetip.Text()), -1)) != 0 {
				for i := 0; i < len(re.FindAllIndex([]byte(linetip.Text()), -1)); i++ {
					find := []int{lines, re.FindAllIndex([]byte(linetip.Text()), -1)[i][0]}
					finds = append(finds, find)
					findlines = append(findlines, linetip.Text())
				}

			}
			lines++
		}
		fmt.Println("----------------Results----------------")
		fmt.Printf("How Many Finds: %v\n", len(finds))
		fmt.Println("Finds:")
		for i := 0; i < len(finds); i++ {
			fmt.Printf("#%v %v:%v  %s\n", i+1, finds[i][0], finds[i][1], findlines[i])
		}
	}
}
