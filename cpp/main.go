package main

import(
  "log"
  "net/http"
  "io/ioutil"
  redis "gopkg.in/redis.v3"
)

var client *redis.Client
 
func main() {
  var err error
  defConf,err=ConfFromYaml(DefConfYaml)
  FATAL("Error reading config file: ")
  appConf,err=ConfFromYamlFile(defConf.Cfg)
  FATAL("Error reading config file: ")
  privateKey,err=ioutil.ReadFile(appConf.Keyfile)
  FATAL("Error reading private key from "+appConf.Keyfile+": ")
  client=redis.NewClient(&redis.Options{Addr:"redis:6379",Password:"",DB: 0,})
  pong,err:=client.Ping().Result()
  FATAL("Error connecting to redis server: ")
  log.Println(pong)
  log.Println(AppName,appConf.Port,AppVersion,"Started")
  http.HandleFunc("/", handleRequest)
  err=http.ListenAndServe(appConf.Port,nil)
  FATAL("Error starting "+AppName+appConf.Port+" "+AppVersion+ ": ")
  log.Println(AppName,appConf.Port,AppVersion,"Stopped")
}
