package main

import (
  "net/http"
  "encoding/json"
)

func resOK(w http.ResponseWriter,res interface{}){
  w.Header().Set("Content-Type","application/json")
  js,_:=json.Marshal(res)
  w.Write(js)
}
