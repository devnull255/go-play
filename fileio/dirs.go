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
  dir,err := os.Open("/Users/devnull")  
  check(err)
  
  defer dir.Close()

  dirnames, err := dir.Readdirnames(0)
  check(err)

  for i := 0;i < len(dirnames); i++ {
    fmt.Println(dirnames[i])
  }
}  
