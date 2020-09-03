package settingsedit

import (
    "os"
    "fmt"
    "regexp"
    "io/ioutil"
    "encoding/json"
)

func OpenJson(filename string) map[string]interface{} {
  jsonFile, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("INFO: Файл открыт")
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)

  re := regexp.MustCompile("(?m)[\r\n]+^.*//.*$")
  lines := re.ReplaceAll([]byte(byteValue), []byte(""))

  result := make(map[string]interface{})
  json.Unmarshal([]byte(lines), &result)
  return result
}

func WriteJson(filename string, result map[string]interface{}) {
  outputfile, err := json.MarshalIndent(result, "", "  ")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("РЕЗУЛЬТАТ:\n", string(outputfile))
  _ = ioutil.WriteFile(filename, outputfile, 0644)
  fmt.Println("INFO: Редактирование и запись в файл завершена")
}

func Cce(filepath string, connstrings ...string) {
  fmt.Println("INFO: Устанавливаем пользовательские параметры в файл:", filepath)
  result := OpenJson(filepath)

  switch {
    case result["cci"] != nil:
      tmp := result["cci"].(map[string]interface{})
      if tmp["connection_string"] != nil {
        tmp["connection_string"] = connstrings[0]
      }
    case result["ConnectionStrings"] != nil:
      tmp := result["ConnectionStrings"].(map[string]interface{})
      if tmp["cci_database"] != nil {
        tmp["cci_database"] = connstrings[0]
      }
      if tmp["cci_hangfire"] != nil {
        tmp["cci_hangfire"] = connstrings[0]
      }
  }
  WriteJson(filepath, result)
}

func Oak(filepath string, connstrings ...string) {
  fmt.Println("INFO: Устанавливаем пользовательские параметры в файл:", filepath)
  result := OpenJson(filepath)
  if result["Data"] != nil {
    tmp := result["Data"].(map[string]interface{})
    tmpconn := tmp["PostgreSql"].(map[string]interface{})
    if tmpconn["ConnectionString"] != nil {
      tmpconn["ConnectionString"] = connstrings[0]
    }
    if tmpconn["HangfireConnectionString"] != nil {
      tmpconn["HangfireConnectionString"] = connstrings[1]
    }
    if tmpconn["UserConnectionString"] != nil {
      tmpconn["UserConnectionString"] = connstrings[2]
    }
  }
  if result["cci"] != nil {
    tmp := result["cci"].(map[string]interface{})
    if tmp["connection_string"] != nil {
      tmp["connection_string"] = connstrings[3]
    }
  }
  WriteJson(filepath, result)
}
