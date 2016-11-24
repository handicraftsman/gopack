package main

import (
  "strings"
  "github.com/handicraftsman/gopack/lib"
)

var procQueueSize = 0
var procQueue [999]map[string]string
var procPkgs = 0
var procPkgsQueue [999]string 
var procProject string
var procVersion string
var procPackage string = "_"
var procBinary string
var procQueueData string = "_"
var procVarCode string

func procParseLine(in string) map[string]string {
  var result = make(map[string]string)
  var split = strings.Split(in, " ")

  if split[0] == "get" {
    result["exists"] = "true"
    result["action"] = split[0]
    result["url"] = split[1]
  } else if split[0] == "project" {
    result["exists"] = "true"
    result["action"] = split[0]
    result["name"] = strings.Join(split[1:len(split)], " ")
  } else if split[0] == "version" {
    result["exists"] = "true"
    result["action"] = "version"
    result["v"] = split[1]
  } else if split[0] == "pkg" {
    result["exists"] = "true"
    result["action"] = split[0]
    result["name"] = split[1]
  } else if split[0] == "binary" {
    result["exists"] = "true"
    result["action"] = split[0]
    result["name"] = split[1]
  } else if split[0] == "queue" {
    result["exists"] = "true"
    result["action"] = "queue"
    result["q"] = split[1]
  } else if split[0] == "git" {
    result["exists"] = "true"
    result["action"] = "git"
    result["repo"] = split[1]
    result["branch"] = split[2]
  } else if split[0] == "var" {
    result["exists"] = "true"
    result["action"] = "var"
    result["name"] = split[1]
    result["type"] = split[2]
    result["val"] = strings.Join(split[3:len(split)], " ")
  } else {
    if strings.Trim(in, " ") == "" {
    } else if in[0] == '#' {
    } else {
      gopack.LogPanic("Unknown action: " + split[0])
    }
  }

  return result
}

func procParseData(in string) {
  lines := strings.Split(in, "\n")

  for i := range lines {
    item := procParseLine(lines[i])
    if item["exists"] == "true" {
      procQueue[procQueueSize] = item
      procQueueSize++
    }
  }
}

func procPre() {
  for i := range procQueue {
    if procQueue[i]["exists"] == "true" {
      if procQueue[i]["action"] == "var" {
        procVarCode = procVarCode +
          "var " + procQueue[i]["name"] + " " +  procQueue[i]["type"] + " = " + procQueue[i]["val"] + "\n"
      }

      if procQueue[i]["action"] == "project" {
        if procProject == "" {
          procProject = procQueue[i]["name"]
        } else {
          gopack.LogPanic("`project` statement can be used only once!")
        }
      }

      if procQueue[i]["action"] == "version" {
        if procVersion == "" {
          procVersion = procQueue[i]["v"]
        } else {
          gopack.LogPanic("`version` statement can be used only once!")
        }
      }

      if procQueue[i]["action"] == "binary" {
        if procBinary == "" {
          procBinary = procQueue[i]["name"]
        } else {
          gopack.LogPanic("`binary` statement can be used only once!")
        }
      }

      if procQueue[i]["action"] == "pkg" {
        if procPackage == "_" {
          procPackage = procQueue[i]["name"]
        } else {
          gopack.LogPanic("`pkg` statement can be used only once!")
        }
      }

      if procQueue[i]["action"] == "queue" {
        if procQueueData == "_" {
          procQueueData = procQueue[i]["q"]
        } else {
          gopack.LogPanic("`queue` statement can be used only once!")
        }
      }
    
      if procQueue[i]["action"] == "get" {
        procPkgsQueue[procPkgs] = procQueue[i]["url"]
      }
    }
  }

  if procProject == "" {
    gopack.LogPanic("You must set project name")
  }
  if procBinary == "" {
    gopack.LogPanic("You must set binary name")
  }
  if procVersion == "" {
    gopack.LogPanic("You must set version")
  }
  if procPackage == "" {
    gopack.LogPanic("You must set package name")
  }

  gopack.LogWrite("# Project: " + procProject + " v" + procVersion + "\n")
  gopack.LogWrite("# Package name: " + procPackage + "\n")
  gopack.LogWrite("# Output file: " + procBinary + "\n")
}
