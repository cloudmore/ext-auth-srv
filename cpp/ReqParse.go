package main

import (
  "log"
  "net/http"
  "errors"
  "encoding/json"
  valid "github.com/asaskevich/govalidator"
)

func parseReq(r *http.Request)(AuthReq,error){
  var req AuthReq
  err:=json.NewDecoder(r.Body).Decode(&req)
  if err!=nil{log.Println("Can't parse request",err)}
  if !valid.IsEmail(req.Email){ return req,errors.New("email is undefined")}
  return req,nil
} 
