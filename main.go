package main

import (
  "flag"
  "io/ioutil"
  "github.com/handicraftsman/gopack/lib"
)

// TODO: Write tests for all statements.
// TODO: Add git-powered package-management.

func main() {
  gopack.LogInit()
  defer gopack.LogStop()
  gopack.LogWrite("# " + Name + " v" + Version + "\n")

  var queue string 
  flag.StringVar(&queue, "q", "", 
    "Sets task-queue. `b` - build; `r` - run; " +
    "`t` - test; `p` - pull packages" + 
    "`i` - install; `g` - generate packdata.go")
  flag.Parse()

  gopack.LogWrite("# Reading Packfile..\n")
  data, err := ioutil.ReadFile("./Packfile")
  if err != nil {
    gopack.LogPanic(err)
  }
  gopack.LogWrite("# Parsing...\n")
  procParseData(string(data))
  gopack.LogWrite("# Preparing environment...\n")
  procPre()

  if (queue == "") && (procQueueData == "") {
    procQueueData = "b"
  } else if queue != "" {
    procQueueData = queue
  }

  for i := range procQueueData {
    switch procQueueData[i] {
    case 'b':
      gopack.LogWrite("# Building...\n")
      procBuild()
      break
    case 'r':
      gopack.LogWrite("# Starting...\n")
      procRun()
      break

    case 't':
      gopack.LogWrite("# Testing...\n")
      procTest()
      break

    case 'p':
      gopack.LogWrite("# Pulling...\n")
      procGet()
      break

    case 'i':
      gopack.LogWrite("# Installing...\n")
      procInstall()
      break

    case 'g':
      gopack.LogWrite("# Generating packdata.go...\n")
      packGenerate()
      break

    default:
      gopack.LogWrite("# Invalid queue-item. Ignoring\n")
      break
    }
  }
}