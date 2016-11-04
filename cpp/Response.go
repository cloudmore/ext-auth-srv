package main

import (
  "io"
  "time"
  "strconv"
  "encoding/json"
  redis "gopkg.in/redis.v3"
  jwt "github.com/dgrijalva/jwt-go"
  xxhash "github.com/OneOfOne/xxhash/native"
)

func respFromToken(token *jwt.Token)(*AuthResp,error){
  tokenStr,err:=token.SignedString(privateKey)
  return &AuthResp{
    token.Claims["id"].(string),
    token.Claims["usr"].(string),
    token.Claims["name"].(string),
    token.Claims["level"].(int),
    token.Claims["exp"].(int64),
    tokenStr,
    tokenStr,
  },err
}

func getResp(req AuthReq)(*AuthResp,error){
  h:=xxhash.New64()
  io.WriteString(h,req.Email)
  id:=h.Sum64()
  uid:=strconv.FormatUint(id,10)
  rv,err:=client.Get(uid).Result()
  if err==redis.Nil{
    usr,_:=json.Marshal(Usr{uid,req.Name,req.Email,appConf.Level})
    client.Set(uid,usr,time.Duration(10)*time.Minute)
  }else if err!=nil{
    return nil,err
  }else {
    var usr Usr
    err=json.Unmarshal([]byte(rv),&usr)
    if usr.Level<=appConf.Level{
      client.Expire(uid,time.Duration(10)*time.Minute)
    }
  }
  exp:=time.Now().Add(time.Hour*1).Unix()
  token:=jwt.New(jwt.GetSigningMethod("RS256"))
  token.Claims["id"]=uid
  token.Claims["exp"]=exp
  token.Claims["usr"]=req.Email
  token.Claims["name"]=req.Name
  token.Claims["level"]=appConf.Level
  return respFromToken(token)
}
