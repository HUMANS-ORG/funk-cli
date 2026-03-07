package commands

import (
 "fmt"
 "os"
 "path/filepath"
)

func FindEmptyFiles(path string) {

 fmt.Println("Scanning for empty files...")

 filepath.Walk(path, func(p string, info os.FileInfo, err error) error {

  if err != nil {
   return nil
  }

  if !info.IsDir() && info.Size() == 0 {
   fmt.Println("Empty file:", p)
  }

  return nil
 })
}