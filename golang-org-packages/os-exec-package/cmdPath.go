package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    path, err := exec.LookPath("go")
    if err != nil {
        log.Fatal("installing go is in your future")
    }
    fmt.Printf("go is available at %s\n", path)
}
