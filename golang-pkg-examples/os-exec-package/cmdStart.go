package main

import (
    "log"
    "os/exec"
)

func main() {
    cmd := exec.Command("sleep", "5")
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Waiting 5 seconds to finished ... ")
    err = cmd.Wait()
    log.Printf("Command finished with error: %v", err)
}
