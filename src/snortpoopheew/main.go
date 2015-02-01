package main

import (
	"fmt"
	"gopkg.in/qml.v0"
	"math/rand"
	"net/url"
	"path"
	"snortpoopheew/src/snortpoopheew/helper"
	//"snortpoopheew/src/snortpoopheew/protocoll"
	"time"
)

var ctrl Control

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	component, err := engine.LoadFile("share/snortpoopheew/main.qml")

	if err != nil {
		panic(err)
	}

	ctrl = Control{Message: "Hello from Go!", StatusText: "initialized..."}

	context := engine.Context()
	context.SetVar("ctrl", &ctrl)

	cfg := SnortConfig{SourceUrl: "", DestinationUrl: "", Hash: "MD5", DestPrepared: false}
	context.SetVar("cfg", &cfg)

	fmt.Println("default hash: " + cfg.Hash)
	fmt.Println("debug: %t", cfg.HasMd5())

	window := component.CreateWindow(nil)

	ctrl.Root = window.Root()

	rand.Seed(time.Now().Unix())

	window.Show()
	window.Wait()
}

type Control struct {
	Root       qml.Object
	Message    string
	StatusText string
}

type SnortConfig struct {
	SourceUrl      string
	DestinationUrl string
	Hash           string
	DestPrepared   bool
}

func (cfg *SnortConfig) SourcePath() string {
	src_url, err := url.Parse(cfg.SourceUrl)
	if err != nil {
		return cfg.SourceUrl
	}
	return src_url.Path
}

func (cfg *SnortConfig) SetDestPrepared() {
	fmt.Println("set dest prepared...")
	cfg.DestPrepared = true
}

func (cfg *SnortConfig) DestinationPath() string {
	dst_url, err := url.Parse(cfg.DestinationUrl)
	if err != nil {
		return cfg.DestinationUrl
	}
	base := path.Base(cfg.SourcePath())
	fmt.Println("Base of path: " + base)
	dest := path.Join(dst_url.Path, base)

	if cfg.DestPrepared != true {
		_, err = helper.PrepareDestPath(dest)
		cfg.SetDestPrepared()
	} else {
		fmt.Println("DestinationPath() called twice...")
	}
	if err != nil {
		fmt.Println("this error should go up: " + err.Error())
	}
	return dest
}

func (cfg *SnortConfig) HasMd5() bool {
	if cfg.Hash == "MD5" {
		return true
	} else {
		return false
	}
}

func (cfg *SnortConfig) HasCrc() bool {
	if cfg.Hash == "CRC32" {
		return true
	} else {
		return false
	}
}

func (ctrl *Control) StartCopy(cfg *SnortConfig) {
	go func() {
		ctrl.StatusText = "starting the process..."
		qml.Changed(ctrl, &ctrl.StatusText)
		helper.CheckDst(cfg.DestinationPath())
		helper.CheckSrc(cfg.SourcePath())
		ctrl.StatusText = "copy to: " + cfg.DestinationPath()
		qml.Changed(ctrl, &ctrl.StatusText)
	}()
}

func SetStatus(ctrl *Control, message string) {
	go func() {
		ctrl.StatusText = message
		qml.Changed(ctrl, &ctrl.StatusText)
	}()
}
