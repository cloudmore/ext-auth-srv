package main

import(
  "log"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

var appConf *AppConf
type AppConf struct {
  Port string `yaml:"port"`
  Level int `yaml:"level"`
  Keyfile string `yaml:"keyfile"`
  Cfg string `yaml:"cfg"`
}

func ConfFromYaml(yml string)(*AppConf,error){
  appConf:=new(AppConf)
  err:=yaml.Unmarshal([]byte(yml),appConf)
  if err!=nil{
    log.Println("Error parsing yml: %s",yml,err)
    return nil,err
  }
  return appConf,nil
}

func ConfFromYamlFile(confFile string)(*AppConf,error){
  ymlBytes,err:=ioutil.ReadFile(confFile);
  if err!=nil{
    log.Println("Error reading config from: %s",confFile,err)
    return nil,err
  }
  appConf:=new(AppConf)
  err=yaml.Unmarshal(ymlBytes,appConf)
  if err!=nil{
    log.Println("Error parsing config from: %s",confFile,err)
    return nil,err
  }
  return appConf,nil
}
