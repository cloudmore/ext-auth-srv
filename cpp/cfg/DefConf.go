package main

import(
  "time"
)

const DefConfYaml=`
#include "defconf.yml"
`
var defConf *AppConf
type ConfFileConf struct{
  FileName string
  RetryCount int
  TimeoutSec time.Duration
}

func defConfFile(defConfFile string)ConfFileConf{
  return ConfFileConf{FileName:defConfFile,RetryCount:-1,TimeoutSec:5}
}
