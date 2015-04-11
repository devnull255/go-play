//return directory names
package main

import (
  "fmt"
  "os"
)


	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  default_dir := os.Getenv("HOME")

  var start_dir string
  var name string

  if len(os.Args) > 1  {
     if os.Args[1] == "-h" || os.Args[1] == "-help" {
        fmt.Println("Usage: dirs.go [dirname]")
        os.Exit(0)
     }
     start_dir = os.Args[1]
  } else {
     start_dir = default_dir
  }
  
  dir,err := os.Open(start_dir)  
  check(err)
  
  defer dir.Close()

  direntries, err := dir.Readdir(0)
  check(err)

  for i := 0;i < len(direntries); i++ {
    if direntries[i].IsDir() {
       name = fmt.Sprintf("[%s]",direntries[i].Name())
    } else {
       name = direntries[i].Name()
    }
    fmt.Println(name)
  }
}  
