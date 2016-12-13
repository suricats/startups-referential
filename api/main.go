package main

import (
    "flag"
    "log"
    "net/http"
    "fmt"
    "github.com/suricats/startups-referential/config"
)

func main() {
    var configArg string
    flag.StringVar(&configArg, "config", "", "toml configuration file")
    flag.Parse()

    configuration := config.ReadConfig(configArg)
    fmt.Printf("%+v", configuration)
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
