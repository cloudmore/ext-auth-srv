package main

import (
  "net/http"
  "encoding/json"
)

type ResError struct {
  Code int `json:"code"`
  Desc string `json:"description"`
  Msg string `json:"message"`
}

func resError(w http.ResponseWriter,code int,msg string){
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(code)
  js,_:=json.Marshal(ResError{code,http.StatusText(code),msg})
  w.Write(js)
}
