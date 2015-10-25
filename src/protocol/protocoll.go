package protocol

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type Protocol struct {
	BasePath string
}

func (p *Protocol) FilePath() string {
	return path.Join(p.BasePath, "protocol.txt")
}

func (p *Protocol) Initialize() (err error) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	now := time.Now()
	logheader := "SnortPoopHeew - Protocol " + now.Format(layout)
	logheaderBytes := []byte(logheader + "\n")

	err = ioutil.WriteFile(p.FilePath(), logheaderBytes, 0644)

	if err != nil {
		fmt.Println("Error creating protokoll-file: ", err)
		return err
	}
	return
}

func (p *Protocol) Success(str string) {
	f := p.getFileHandle()
	_, err := f.WriteString("SUCCESS: " + str + "\n")
	if err != nil {
		fmt.Println("Error writing protocol: ", err)
		panic(err)
	}
	defer f.Close()
}

func (p *Protocol) Failure(str string) {
	f := p.getFileHandle()
	_, err := f.WriteString("ERROR: " + str + "\n")
	if err != nil {
		fmt.Println("Error writing protocol: ", err)
	}
	defer f.Close()
}

func (p *Protocol) getFileHandle() (file *os.File) {
	f, err := os.OpenFile(p.FilePath(), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error - could not open Protocol-file: ", err)
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
