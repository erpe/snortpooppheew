package main

import (
	"./gopkg.in/qml.v1"
	"./helper"
	"./protocol"
	"./worker"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var ctrl Control

func main() {
	err := qml.Run(run)
	if err != nil {
		panic(err)
	}
}

func run() error {
	//qml.Init(nil)
	engine := qml.NewEngine()
	component, err := engine.LoadFile("../share/snortpoopheew/main.qml")

	if err != nil {
		return err
	}

	ctrl = Control{Message: "Hello from Go!", StatusText: "initialized..."}

	context := engine.Context()
	context.SetVar("ctrl", &ctrl)

	cfg := worker.SnortConfig{SourceUrl: "", DestinationUrl: "", Hash: "MD5", DestPrepared: false}
	context.SetVar("cfg", &cfg)

	fmt.Println("default hash: " + cfg.Hash)
	fmt.Println("debug: %t", cfg.HasMd5())

	window := component.CreateWindow(nil)

	ctrl.Root = window.Root()

	rand.Seed(time.Now().Unix())

	window.Show()
	window.Wait()
	return err
}

type Control struct {
	Root       qml.Object
	Message    string
	StatusText string
}

func (ctrl *Control) StartCopy(cfg *worker.SnortConfig) {
	go func() {
		SetStatus(ctrl, "starting the process...")

		helper.CheckDst(cfg.DestinationPath())
		helper.CheckSrc(cfg.SourcePath())

		msg := "copy to: " + cfg.DestinationPath()
		SetStatus(ctrl, msg)

		proto := protocol.Protocol{cfg.DestinationPath()}
		err := proto.Initialize()

		if err != nil {
			fmt.Println("initialize protocoll: " + err.Error())
			panic(err)
		}

		cpm := worker.CopyMachine{cfg.SourcePath(), cfg.DestinationPath(), cfg.Hash, &proto, 0}

		err = cpm.Process()

		if err != nil {
			fmt.Println("ERRR: " + err.Error())
		}
		if cpm.ErrorCount > 0 {
			c := strconv.Itoa(cpm.ErrorCount)
			msg = "Finished with errors: " + c + " Errors encounted... see log for details"
		} else {
			msg = "Finished successfully with no errors... "
		}

		SetStatus(ctrl, msg)
		ctrl.Root.ObjectByName("StartBtn").Set("enabled", true)
	}()
}

func SetStatus(ctrl *Control, message string) {
	go func() {
		ctrl.StatusText = message
		qml.Changed(ctrl, &ctrl.StatusText)
	}()
}
