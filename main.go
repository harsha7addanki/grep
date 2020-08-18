package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	} else if Args[0] == "FINDALL" {
		fmt.Println("----------------Results----------------")
		findAll(Args[1], Args[2])
	}
}

func findAll(arg1 string, arg2 string) {
	dir, err := os.Open(arg1)
	if err != nil {
		//fmt.Println(Red + "[ERROR]Failed To Open Directory: " + err + Reset)
		panic(err)
	}

	defer dir.Close()
	list, _ := dir.Readdirnames(0)
	finds := [][]int{}
	findlines := []string{}
	findfiles := []string{}
	for _, name := range list {
		//fmt.Printf("\nprocessing file " + name + "\n")

		fileInfo, err := os.Stat(arg1 + "/" + name)
		if err != nil {
			log.Fatal(err)
		}

		if fileInfo.IsDir() {
			//fmt.Println("Preocessing Folder")
			findAll(arg1+"/"+name, arg2)
		}

		file, err := os.Open(arg1 + "/" + name)
		ftype, err := GetFileContentType(file)
		if err != nil {
			//log.Fatal(err)
			continue
		}
		if ftype == "application/octet-stream" {
			//fmt.Println("Skipping binary file")
			continue
		}
		linetip := bufio.NewScanner(file)
		var lines int = 0
		re := regexp.MustCompile(arg2)
		cont := false
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
		if cont {
			findfiles = append(findfiles, name)
		}
	}

	for i, j := 0, 0; i < len(finds); i, j = i+1, j+1 {
		fmt.Printf("#%v %v:%v  %s file: %s\n", i+1, finds[i][0], finds[i][1], findlines[i], findfiles[j])
	}

}

/**
This is comment
**/
func GetFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
