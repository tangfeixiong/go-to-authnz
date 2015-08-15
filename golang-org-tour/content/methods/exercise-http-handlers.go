// +build OMIT

package main

import (
        "fmt"
	"log"
	"net/http"
        "os"
        //"time"
)

var s http.Server

func main() {
	// your http.Handle calls here
        var handleHome String = String("index")
        var handleRest Struct = Struct{"welcome", "see you", "Gophers!"}

        http.Handle("/index", handleHome)
        http.Handle("/service/", handleRest)

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
        /*
        s = http.Server {
            Addr          : ":4000",
            Handler       : handleRest,
            ReadTimeout   : 10 * time.Second,
            WriteTimeout  : 10 * time.Second,
            MaxHeaderBytes: 1 << 20,   
        }

        log.Fatal(s.ListenAndServe())
        */
}

type String string

type Struct struct {
    Greeting string
    Goodbye  string
    Who      string
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(os.Stdout, r.URL.Path)
    fmt.Fprintf(w, "Hello! %q\n", r.UserAgent())
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(os.Stdout, r.URL.Path)
    switch r.URL.Path {
    case "/service/greeting":
        fmt.Fprint(w, h.Greeting)
    case "/service/goodbye":
        fmt.Fprint(w, h.Goodbye)
    case "/service/who":
        fmt.Fprint(w, h.Who)
    default:
        http.NotFound(w, r)
    }
}
