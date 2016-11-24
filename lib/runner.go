package gopack

// #include <stdlib.h>
// int gopack_run(char* cmd) {
//   return system(cmd);
// }
// typedef char* (*strFunc) ();
import "C"
import (
  "strconv"
)

func Run(cmd string) {
  _code := C.gopack_run(C.CString(cmd))
  code := int(_code)

  if code != 0 {
    LogPanic("Process errored with code [" + strconv.Itoa(code) + "]")
  }
}