package main

import(
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
  jwt "github.com/dgrijalva/jwt-go"
)

func myLookupKey(kid interface{})(interface{},error){
  log.Println("Server Stop")
  verifyBytes,err:=ioutil.ReadFile("bin/auth-srv.pub")
  FATAL("Error reading pub key file: ")
  return jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
}

func validate(w http.ResponseWriter,r *http.Request)(){
  _,_=jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{},error){
    if _,ok:=token.Method.(*jwt.SigningMethodRSA); !ok{
      return nil,fmt.Errorf("Unexpected signing method: %v",token.Header["alg"])
    }
    return myLookupKey(token.Header["kid"])
  })
}
