package main

type AuthResp struct {
  Id string `json:"id"`
  Usr string `json:"email"`
  Name string `json:"name"`
  Level int `json:"level"`
  Exp int64 `json:"exp"`
  Token string `json:"token"`
  Refresh string `json:"refresh"`
}
