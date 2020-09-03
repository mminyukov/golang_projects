package configoak

import (
  "os"
  "fmt"
  "encoding/json"
)

type Config struct {
  Oak struct {
    Stand_name string `json"stand_name"`
    Port_site string `json"port_site"`
    Target_service_name string `json"target_service_name"`
    Connection_string_user string `json"connection_string_user"`
    Connection_string_db string `json"connection_string_db"`
    Connection_string_hangfire string `json"connection_string_hangfire"`
    Connection_string_cci string `json"connection_string_cci"`
    Ss_url string `json"ss_url"`
    Dbupdater_zip_name string `json"dbupdater_zip_name"`
    Main_zip_name string `json"main_zip_name"`
    Use_scheduler bool `json"use_scheduler"`
    Port_scheduler string `json"port_scheduler"`
    Target_scheduler_name string `json"target_scheduler_name"`
    Connection_string_scheduler string `json"connection_string_scheduler"`
    Scheduler_zip_name string `json"scheduler_zip_name"`
    Recreate_db bool `json"recreate_db"`
    Install_cce string `json"install_cce"`
  } `json:"oak"`
  Mainconf Mainconf `json:"main"`
}

type Mainconf struct {
  Prefix_directory string `json:"prefix_directory"`
}

func LoadJson(filename string) (Config, error) {
  var config Config
  jsonFile, err := os.Open(filename)
  defer jsonFile.Close()
  if err != nil {
    return config, err
    fmt.Println(err)
  }
  jsonParser := json.NewDecoder(jsonFile)
  err = jsonParser.Decode(&config)
  return config, err
}
