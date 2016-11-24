package gopack

import (
  "os"
)

func FileExists(file string) (bool, error) {
  _, err := os.Stat(file)
  
  if err == nil { 
    return true, nil
  }

  if os.IsNotExist(err) { 
    return false, nil
  }
  
  return true, err
}