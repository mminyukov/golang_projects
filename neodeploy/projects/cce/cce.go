package cce

import (
  "os"
  "fmt"
  "path"
  "neodeploy/service/manage"
  "neodeploy/service/unzip"
  "neodeploy/service/recreatedb"
  "neodeploy/service/settingsedit"
  "neodeploy/service/templates/unit"
  "neodeploy/service/configparser/cce"
)

func Install(filename string) {
  config, _ := configcce.LoadJson(filename)
  var prefix_directory string = "/usr/local/share"
  var target_directory string = path.Join(prefix_directory,"cce",config.Cce.Stand_name)
  var target_directory_site string = path.Join(target_directory,"JobStarter")
  var target_directory_dbupdater string = path.Join(target_directory,"DbUpdater")
  var site_service_name string = "cce." + config.Cce.Stand_name + ".service"
  var site_service_file string = path.Join("/etc/systemd/system/", site_service_name)

  manage.CheckFile(config.Cce.Main_zip_name)
  if _, err := os.Stat(site_service_file); err == nil {
    manage.Action("stop",site_service_name)
    manage.Action("disable",site_service_name)
  }

  if _, err := os.Stat(target_directory_site); err == nil {
    fmt.Println("INFO: Очищаем директорию:",target_directory)
    os.RemoveAll(target_directory)
  }

  uz := unzip.New(config.Cce.Main_zip_name, target_directory)
  err := uz.Extract()
  if err != nil {
    fmt.Println(err)
  }

  settingsedit.Cce(path.Join(target_directory_site,"settings.json"),config.Cce.Connection_string)

  if config.Cce.Recreate_db == true {
    settingsedit.Cce(path.Join(target_directory_dbupdater,"settings.json"),config.Cce.Connection_string)
    os.Chmod(path.Join(target_directory_dbupdater,"Tenax.CCI.DbUpdater.dll"), 0777)
    fmt.Println("WARN: База данных будет пересоздана")
    recreatedb.Cce(target_directory_dbupdater)
  }

  createunit.Createunit("cce",config.Cce.Stand_name,target_directory_site,config.Cce.Target_service_name,site_service_file)
  os.Chmod("some-filename", 0777)
  manage.Action("daemon-reload",site_service_name)
  manage.Action("enable",site_service_name)
  manage.Action("start",site_service_name)
  fmt.Println("INFO: Закончили")
}