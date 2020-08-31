package oak

import (
  "fmt"
  "neodeploy/service/configparser/oak"
)

func Install(filename string) {
  config, _ := configoak.LoadJson(filename)
  fmt.Println(config.Oak)
}