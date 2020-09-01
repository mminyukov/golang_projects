package main

import (
  "os"
  "fmt"
  "neodeploy/projects/cce"
  "neodeploy/projects/oak"
  "neodeploy/service/manage"
)

func main() {
  var settingsdeploy string = "settingsdeploy.json"
  manage.CheckFile(settingsdeploy)
  if len(os.Args) == 1 {
    fmt.Println("ERROR: Требуется параметр для запуска")
    os.Exit(0)
  } else if len(os.Args) > 2 {
    fmt.Println("ERROR: Должен быть один параметр")
  } else {
    switch {
      case os.Args[1] == "oak":
        oak.Install(settingsdeploy)
      case os.Args[1] == "cce":
        cce.Install(settingsdeploy)
    }
  }
}