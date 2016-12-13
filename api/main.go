package main

import (
    "github.com/suricats/startups-referential/config"
    "github.com/suricats/startups-referential/db"
    "flag"
    "log"
    "net/http"
)

func main() {
    var configArg string
    flag.StringVar(&configArg, "config", "", "toml configuration file")
    flag.Parse()

    configuration := config.ReadConfig(configArg)
    connection := db.GetConnection(configuration)
    defer connection.Close()
    db.ListStartups(configuration, connection)

    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
