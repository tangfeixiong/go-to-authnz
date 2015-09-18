package main
import (
    "fmt"
    "os"
)

func main() {
    gopath := os.Getenv("GOPATH")
    if gopath == "" {
        fmt.Printf("GOPATH is not exported")
        return
    }
    fmt.Println(gopath)
    
    // undefined in version 1.2.1
    gopath, ok := os.LookupEnv("NOTEXPORTED")
    if !ok {
        fmt.Println("NOTEXPORTED is not exported")
        return
    }
    fmt.Println(gopath)
}