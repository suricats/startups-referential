package main

import (
    "flag"
    "log"
    "net/http"
    "fmt"
)

func main() {
    var configArg string
    flag.StringVar(&configArg, "config", "", "toml configuration file")
    flag.Parse()

    config := ReadConfig(configArg)
    fmt.Printf("%+v", config)
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
