package gopack

import (
  "fmt"
  "log"
  "os"
)

var logFile *os.File
var logRunning = false

func LogInit() {
  exists, err := FileExists("./packer.log")
  if err != nil {
    log.Panic(err)
  }

  if !exists {
    os.Remove("./packer.log")
  }

  logFile, err = os.Create("./packer.log")
  if err != nil {
    log.Panic(err)
  }
}

func LogWrite(msg interface{}) {
  fmt.Printf(msg.(string))
  logFile.WriteString(msg.(string))
}

func LogPanic(msg interface{}) {
  log.Panic(msg)
  logFile.WriteString(msg.(string))
}

func LogStop() {
  logFile.Close()
}