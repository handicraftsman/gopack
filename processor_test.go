package main;

import (
  "fmt"
  "testing"
)

func testCompareStringMaps(map1, map2 map[string]string) bool {
  out := true

  for i, value := range map2 {
    if map1[i] != value {
      out = false
    }
  }

  if !out {
    fmt.Println("\n Got:", map2, "\n Expected: ", map1)
  }

  return out
}

func TestGetStatement(t *testing.T) {
  var test = make(map[string]string)
  test["action"] = "get"
  test["url"] = "github.com/go-bottle/go-bottle" 
  test["exists"] = "true"

  var test2 = procParseLine("get github.com/go-bottle/go-bottle")

  if !testCompareStringMaps(test, test2) {
    t.Fail()
  }
}

func TestProjectStatement(t *testing.T) {
  var test = make(map[string]string)
  test["action"] = "project"
  test["name"] = "Sample Project. Hello World!"
  test["exists"] = "true"

  var test2 = procParseLine("project Sample Project. Hello World!")

  if !testCompareStringMaps(test, test2) {
    t.Fail()
  }
}

func TestBinaryStatement(t *testing.T) {
  var test = make(map[string]string)
  test["action"] = "binary"
  test["name"] = "go_packer"
  test["exists"] = "true"

  var test2 = procParseLine("binary go_packer")

  if !testCompareStringMaps(test, test2) {
    t.Fail()
  }
}