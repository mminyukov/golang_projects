package oak

import (
  "os"
  "fmt"
  "path"
  "neodeploy/service/unzip"
  "neodeploy/service/manage"
  "neodeploy/service/recreatedb"
  "neodeploy/service/settingsedit"
  "neodeploy/service/templates/unit"
  "neodeploy/service/configparser/oak"
)

func Install(filename string) {
  config, _ := configoak.LoadJson(filename)
  var target_directory string = path.Join(config.Mainconf.Prefix_directory,"oak",config.Oak.Stand_name)
  var target_directory_site string = path.Join(target_directory,"Site")
  var target_directory_dbupdater string = path.Join(target_directory,"DbUpdater")
  var site_service_name string = "oak." + config.Oak.Stand_name + ".service"
  var site_service_file string = path.Join("/etc/systemd/system/", site_service_name)
  var site_user_name string = "www-data"
  var site_run_directory string = path.Join("/var/lib","oak",config.Oak.Stand_name)

  var target_directory_scheduler string = path.Join(target_directory,"Scheduler")
  var scheduler_service_name string = "oak." + config.Oak.Stand_name + ".shceduler.service"
  var scheduler_service_file string = path.Join("/etc/systemd/system/",scheduler_service_name)
  var scheduler_run_directory string = "/var/lib/Scheduler"

  manage.CheckFile(config.Oak.Main_zip_name)
  manage.CheckFile(config.Oak.Dbupdater_zip_name)
  if config.Oak.Use_scheduler == true {
    manage.CheckFile(config.Oak.Scheduler_zip_name)
  }
  if _, err := os.Stat(site_service_file); err == nil {
    manage.Action("stop",site_service_name)
    manage.Action("disable",site_service_name)
  }

  manage.RemoveDir(target_directory_site)
  uz := unzip.New(config.Oak.Main_zip_name, target_directory_site)
  err := uz.Extract()
  if err != nil {
    fmt.Println(err)
  }

  settingsedit.Oak(path.Join(target_directory_site,"appsettings.json"), "", config.Oak.Connection_string_hangfire, config.Oak.Connection_string_user, config.Oak.Connection_string_cci)

  if config.Oak.Recreate_db == true {
    manage.RemoveDir(target_directory_dbupdater)
    uz := unzip.New(config.Oak.Dbupdater_zip_name, target_directory_dbupdater)
    err := uz.Extract()
    if err != nil {
      fmt.Println(err)
    }
    settingsedit.Oak(path.Join(target_directory_dbupdater,"appsettings.json"),config.Oak.Connection_string_db)
    os.Chmod(path.Join(target_directory_dbupdater,"Neolant.OAK.DbUpdater"), 0777)
    fmt.Println("WARN: База данных будет пересоздана")
    recreatedb.Oak(target_directory_dbupdater)
  }

  if config.Oak.Use_scheduler == true {
    manage.RemoveDir(target_directory_scheduler)
    uz := unzip.New(config.Oak.Scheduler_zip_name, target_directory_scheduler)
    err := uz.Extract()
    if err != nil {
      fmt.Println(err)
    }
    createunit.Createunit("oak.scheduler",scheduler_service_file,config.Oak.Stand_name,target_directory_scheduler,config.Oak.Target_scheduler_name,site_user_name,scheduler_run_directory,config.Oak.Port_scheduler)
    os.Chmod(path.Join(target_directory_scheduler,config.Oak.Target_scheduler_name), 0777)
    manage.Action("daemon-reload",scheduler_service_name)
    manage.Action("enable",scheduler_service_name)
    manage.Action("start",scheduler_service_name)
  }

  createunit.Createunit("oak",site_service_file,config.Oak.Stand_name,target_directory_site,config.Oak.Target_service_name,site_user_name,site_run_directory,config.Oak.Port_site)
  os.Chmod(path.Join(target_directory_site,config.Oak.Target_service_name), 0777)
  manage.Action("daemon-reload",site_service_name)
  manage.Action("enable",site_service_name)
  manage.Action("start",site_service_name)
  fmt.Println("INFO: Закончили")
}