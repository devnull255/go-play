package main

import (
  "crypto/rand"
  "crypto/rsa"
  "fmt"
)

func main() {
   privkey,err := rsa.GenerateKey(rand.Reader,256)
   if err != nil {
      panic(err)
   }
   fmt.Println(privkey.Public())
}
