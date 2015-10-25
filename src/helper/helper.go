package helper

import (
	"fmt"
	"os"
)

func PrepareDestPath(path string) (string, error) {
	ret, _ := exists(path)

	if ret != false {
		message := "destination path exists already... "
		err := fmt.Errorf(message)
		printErrAndExit(message)
		return "", err
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("Error - could not create protokoll-directory: ", path)
		return "", err
	}

	return path, nil
}

func CheckSrc(url string) bool {
	var ret, _ = exists(url)
	if ret == false {
		message := "No such sourcedir: " + url
		printErrAndExit(message)
		return false
	}
	return true
}

func CheckDst(url string) bool {
	var ret, _ = exists(url)
	if ret == false {
		message := "No such destinationdir: " + url
		printErrAndExit(message)
		return false
	}
	return true
}

func exists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func printErrAndExit(msg string) {
	fmt.Println(msg)
	//os.Exit(1)
}
