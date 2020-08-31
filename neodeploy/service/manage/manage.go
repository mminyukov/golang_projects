package main

import (
    "os/exec"
    "fmt"
    "os"
)

func main() {
    cmd := exec.Command("systemctl", "check", "sshd")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Cannot find process")
        os.Exit(1)
    }
    fmt.Printf("Status is: %s", string(out))
    fmt.Println("Starting Role")
}