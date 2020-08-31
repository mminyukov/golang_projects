package cce

import (
  "os"
  "fmt"
  "neodeploy/service/configparser/cce"
)

func Install(filename string) {
  config, _ := configcce.LoadJson(filename)
  if _, err := os.Stat(config.Cce.Main_zip_name); os.IsNotExist(err) {
    fmt.Println("Файл не найден")
    os.Exit(1)
  } else {
    
  }
  fmt.Println(config.Cce)
}