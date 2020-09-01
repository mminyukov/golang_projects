package manage

import (
    "os/exec"
    "fmt"
    "os"
)

func Action(action string, service string) {
  cmd := exec.Command("systemctl", "check", service)
  out, _ := cmd.CombinedOutput()
/*
  if err != nil {
    if exitErr, ok := err.(*exec.ExitError); ok {
      fmt.Printf("Сервис отключен с ненулевым кодом: %v\n", exitErr)
    } else {
      fmt.Printf("Ошибка запуска systemctl: %v", err)
      os.Exit(1)
    }
  }
*/
  fmt.Printf("INFO: Статус сервиса %s: %s", service, string(out))
  fmt.Println("WARN: Выполняем действие:", action)
  if action == "daemon-reload" {
  cmd := exec.Command("systemctl", action)
  cmd.Run()
  } else {
    cmd := exec.Command("systemctl", action, service)
    cmd.Run()
  }
}

func CheckFile(filename string) {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    fmt.Println("ERROR: Файл не найден:", filename)
    os.Exit(1)
  }
}