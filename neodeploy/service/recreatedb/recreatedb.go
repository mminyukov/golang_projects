package recreatedb

import (
  "os"
  "fmt"
  "log"
  "os/exec"
)

func Cce(workdir string) {
  currentdir, _ := os.Getwd()
  os.Chdir(workdir)
  cmd := exec.Command("dotnet", "Tenax.CCI.DbUpdater.dll", "-r", "-s")
  stdoutStderr, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", stdoutStderr)
  os.Chdir(currentdir)
}

func Oak(workdir string) {
  currentdir, _ := os.Getwd()
  os.Chdir(workdir)
  cmd := exec.Command("./Neolant.OAK.DbUpdater", "-r", "-f")
  stdoutStderr, err := cmd.CombinedOutput()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", stdoutStderr)
  os.Chdir(currentdir)
}