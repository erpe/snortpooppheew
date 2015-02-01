package protocoll

import (
	"fmt"
	"time"
	"io/ioutil"
	"os"
)

var filename string

func Initialize(dest string) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	createLogDir(dest)
	now := time.Now()
	logheader := "File-Copy - Protocol " + now.Format(layout)
	filename = dest + "/protocol.txt"
	logheaderBytes := []byte(logheader + "\n")

	err := ioutil.WriteFile(filename, logheaderBytes, 0644)

	if err != nil {
		fmt.Println("Error creating protokoll-file: ", err)
	}
}

func Success(str string) {
	f := getFileHandle()
	_, err := f.WriteString("SUCCESS: " + str + "\n")
	if err != nil {
		fmt.Println("Error writing protocol: ", err)
		panic(err)
	}
	defer f.Close()
}

func Failure(str string) {
	f := getFileHandle()
	_, err := f.WriteString("ERROR: " + str + "\n")
	if err != nil {
		fmt.Println("Error writing protocol: ", err)
	}
	defer f.Close()
}

func getFileHandle() (file *os.File) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error - could not open Protokoll-file: ", err)
		panic(err)
	}
	file = f
	return file
}

func createLogDir(dest string) {
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		fmt.Println("Error - could not create protokoll-directory: ", dest)
		panic(err)
	}
}
