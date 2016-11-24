package main

import "github.com/handicraftsman/gopack/lib"

func procBuild() {
  gopack.Run("go build -v -o " + procBinary)
}

func procTest() {
  gopack.Run("go test -v")
}

func procRun() {
  gopack.Run("./" + procBinary)
}

func procGet() {
  for i := range procPkgsQueue {
    if procPkgsQueue[i] != "" {
      gopack.Run("go get -u -v " + procPkgsQueue[i])
    }
  }
}

/**
  ПРОЧИТАТЬ ПЕРЕД РЕАЛИЗАЦИЕЙ:

  О vendor:
  https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit#

  Этот документ сообщает, что, скажем, можно склонировать go-bottle в
  ./vendor/github.com/go-bottle/go-bottle/ и он будет доступен для импорта как
  "github.com/go-bottle/go-bottle". На самом же деле это выражение будет расширяться до
  "github.com/handicraftsman/gopack/vendor/go-bottle/go-bottle"

  Страница ниже подтверждает сказанное:
  https://blog.gopheracademy.com/advent-2015/vendor-folder/

  Также из второго документа можно извлечь, что приоритет у папки vendor 
  выше, чем у $GOPATH

  Таким образом, реализовать git-powered систему зависимостей не составляет труда.
**/


func procInstall() {
  gopack.Run("go install")
}