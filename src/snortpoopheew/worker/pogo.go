package worker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"snortpoopheew/src/snortpoopheew/protocol"
	"strconv"
)

type CopyError struct {
	What string
}

func (e *CopyError) Error() string {
	return e.What
}

type CopyMachine struct {
	SrcDir     string
	DstDir     string
	Hash       string
	Protocol   *protocol.Protocol
	ErrorCount int
}

var machine *CopyMachine

func (cm *CopyMachine) Initialize(source string, destination string, hash string) *CopyMachine {
	cm.SrcDir = source
	cm.DstDir = destination
	cm.Hash = hash
	return cm
}

func (cm *CopyMachine) Process() (err error) {
	machine = cm
	err = CopyDir(cm.SrcDir, cm.DstDir)
	if err != nil {
		return err
	}
	return
}

func CopyDir(source string, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		msg := "Source is not a directory: " + source
		return &CopyError{msg}
	}

	_, err = os.Open(dest)

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {
		sfp := path.Join(source, entry.Name())
		dfp := path.Join(dest, entry.Name())
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return
}

func CopyFile(source string, dest string) (err error) {

	sf, err := os.Open(source)

	if err != nil {
		return err
	}

	defer sf.Close()

	df, err := os.Create(dest)

	if err != nil {
		return err
	}

	defer df.Close()

	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

		var destSum, sourceSum string
		if machine.Hash == "MD5" {
			destSum = checkMd5(sf)
			sourceSum = checkMd5(df)
		}

		if machine.Hash == "CRC32" {
			destSum = checkCrc32(source)
			sourceSum = checkCrc32(dest)
		}

		if destSum == sourceSum {
			go machine.Protocol.Success(dest + " : " + sourceSum + " : GOOD")
			go fmt.Print("+")
		} else {
			machine.ErrorCount += 1
			go machine.Protocol.Failure(source + " : " + sourceSum + " : MISMATCH")
			go fmt.Print("E")
		}
	}
	return
}

func checkMd5(file io.Reader) (sum string) {
	h := md5.New()
	io.Copy(h, file)
	sum = hex.EncodeToString(h.Sum(nil))
	return sum
}

func checkCrc32(fileName string) (sum string) {
	h := crc32.NewIEEE()
	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic("(crc32) Error reading: " + fileName)
	} else {
		h.Write(b)
		sum = strconv.FormatUint(uint64(h.Sum32()), 10)
		return sum
	}

	return
}
