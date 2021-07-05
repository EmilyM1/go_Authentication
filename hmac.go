package main

import (
  "fmt"
  "crypto/hmac"
  "crypto/sha512"

)

var key = []byte{}
var msg string


func main() {
  //make the key, just for example purposes for key creation, not actual key
  for i := 1; i < 64; i++{
    key = append(key, byte(i))
    }
    msg = "hello"
    fmt.Println(signMsg([]byte(msg)))

  }

//use hmac to sign a message
func signMsg(msg []byte) ([]byte){
  // takes in algorithm as first parameter, 2nd parameter is a key which is a byte slice
  h := hmac.New(sha512.New, key)
  _, err := h.Write(msg)
  if err != nil{
    return nil
  }
  signedwithkey := h.Sum(nil)
  return signedwithkey
}

// can use hmac equal to verify match
