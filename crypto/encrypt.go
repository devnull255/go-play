package main

import (
  "crypto/rand"
  "crypto/rsa"
  "crypto/md5"
  "fmt"
)

func main() {
   privkey,err := rsa.GenerateKey(rand.Reader,512)
   if err != nil {
      panic(err)
   }
   str := []byte("Hello I have a secret.")
   fmt.Println(str)
   md5hash := md5.New()
   label := []byte("foo")
   var publickey *rsa.PublicKey
   publickey = &privkey.PublicKey
   
   encrypted_data,err:= rsa.EncryptOAEP(md5hash,rand.Reader,publickey,str,label)
   if err != nil {
     panic(err)
   }
   fmt.Println("Encrypted text:",string(encrypted_data))
   decrypted_data,err := rsa.DecryptOAEP(md5hash,rand.Reader,privkey,encrypted_data,label) 
   if err != nil {
     panic(err)
   }
   fmt.Println("Decrypted text: ",string(decrypted_data))
}
