package main

import (
    "fmt"
    "bytes"
    "log"
    "os/exec"
    "strings"
)

func main() {
    cmd := exec.Command("echo", "this command")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("result: %q\n", out.String())
}
