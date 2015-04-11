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

  if len(os.Args) > 1  {
     start_dir = os.Args[1]
  } else {
     start_dir = default_dir
  }
  
  dir,err := os.Open(start_dir)  
  check(err)
  
  defer dir.Close()

  dirnames, err := dir.Readdirnames(0)
  check(err)

  for i := 0;i < len(dirnames); i++ {
    fmt.Println(dirnames[i])
  }
}  
