package createunit

import (
  "os"
  "fmt"
  "text/template"
)

type Vars struct {
  Project_name, Stand_name, Target_directory_site, Target_service_name string
}

func Createunit(project_name string, stand_name string, target_directory_site string ,target_service_name string, site_service_file string) {
  v := Vars{project_name,stand_name,target_directory_site,target_service_name}
  tpl, _ := template.ParseFiles("service/templates/unit/unit.txt")
  f, err := os.OpenFile(site_service_file, os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer f.Close()

  tpl.Execute(f, v)
}