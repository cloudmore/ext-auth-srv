package main

type AuthReq struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  Provider string `json:"_provider"`
}
