package main

import(
  "net/http"
)

var privateKey []byte

func handleRequest(w http.ResponseWriter,r *http.Request){
  if r.Method=="OPTIONS"{var a struct{};resOK(w,a);return}
  if r.Method!="POST"{resError(w,http.StatusMethodNotAllowed ,"unsupported request");return}
  req,err:=parseReq(r)
  if err!=nil{resError(w,http.StatusBadRequest,err.Error());return}
  resp,err:=getResp(req)
  if err!=nil{resError(w,http.StatusInternalServerError,"can't generate response");return}
  resOK(w,resp)
} 
