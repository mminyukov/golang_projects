package createunit

import (
  "os"
  "fmt"
  "text/template"
)

func Createunit(project_name string, site_service_file string, vars ...string) {
  var v map[string]string
  tpl := template.New("")
  switch {
    case project_name == "cce":
      v = map[string]string{"Project_name": project_name, "Stand_name": vars[0], "Target_directory_site": vars[1], "Target_service_name": vars[2]}
      tpl, _ = template.ParseFiles("service/templates/unit/unitcce.txt")
    case project_name == "oak":
      v = map[string]string{"Project_name": project_name, "Stand_name": vars[0], "Target_directory_site": vars[1], "Target_service_name": vars[2], "Site_user_name": vars[3], "Site_run_directory": vars[4], "Port_site": vars[5]}
      tpl, _ = template.ParseFiles("service/templates/unit/unitoak.txt")
    case project_name == "oak.scheduler":
      v = map[string]string{"Project_name": project_name, "Stand_name": vars[0], "Target_directory_site": vars[1], "Target_service_name": vars[2], "Site_user_name": vars[3], "Site_run_directory": vars[4], "Port_site": vars[5]}
      tpl, _ = template.ParseFiles("service/templates/unit/unitoak.txt")
  }
  if _, err := os.Stat(site_service_file); err != nil {
    os.Create(site_service_file)
  }
  f, err := os.OpenFile(site_service_file, os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer f.Close()
  tpl.Execute(f, v)

}
