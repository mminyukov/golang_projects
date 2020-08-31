package configcce

import (
  "os"
  "fmt"
  "encoding/json"
)

type Config struct {
  Cce struct {
    Stand_name string `json:"stand_name"`
    Target_service_name string `json:"target_service_name"`
    Recreate_db bool `json:"recreate_db"`
    Connection_string string `json:"connection_string"`
    Main_zip_name string `json:"main_zip_name"`
  } `json:"cce"`
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
