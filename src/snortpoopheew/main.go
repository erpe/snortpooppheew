package main



import (
        "fmt"
        "gopkg.in/qml.v0"
        "math/rand"
        "time"
)


func main() {
        qml.Init(nil)
        engine := qml.NewEngine()
        component, err := engine.LoadFile("share/snortpoopheew/main.qml")
        if err != nil {
                panic(err)
        }

        ctrl := Control{Message: "Hello from Go!", StatusText: "initialized..."}

        context := engine.Context()
        context.SetVar("ctrl", &ctrl)

        cfg := SnortConfig{ Source: "", Destination: "", Hash: "MD5"}
        context.SetVar("cfg", &cfg)

        fmt.Println("default hash: " + cfg.Hash)
        fmt.Println("debug: %t", cfg.HasMd5() )
        window := component.CreateWindow(nil)

        ctrl.Root = window.Root()

        rand.Seed(time.Now().Unix())

        window.Show()
        window.Wait()
}

type Control struct {
        Root    qml.Object
        Message string
        StatusText string
}

type SnortConfig struct {
  Source string
  Destination string
  Hash string
}

func (cfg *SnortConfig) HasMd5() bool {
  if (cfg.Hash == "MD5") {
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

func (ctrl *Control) Hello() {
    go func() {
      ctrl.Message = "Hello from Go Again!"
      qml.Changed(ctrl, &ctrl.Message)
    }()
}

func (ctrl *Control) SourceDir(txt string) {
  go func() {
    ctrl.StatusText = txt
    qml.Changed(ctrl, &ctrl.StatusText)
  }()
}

func (ctrl *Control) DestinationDir(txt string) {
  go func() {
    ctrl.StatusText = txt
    qml.Changed(ctrl, &ctrl.StatusText)
  }()
}

func (ctrl *Control) StartCopy() {
  go func() {
    ctrl.StatusText = "starting the process..."
    qml.Changed(ctrl, &ctrl.StatusText)
  }()
}
