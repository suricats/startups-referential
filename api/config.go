package main

import (
    "os"
    "log"
    "github.com/BurntSushi/toml"
)

type DatabaseConfig struct {
    User string
    Password string
    Host string
    Port int
}

type Configuration struct {
    DatabaseConfig DatabaseConfig `toml:"database"`
}

func ReadConfig(configfile string) Configuration {
    _, err := os.Stat(configfile)
    if err != nil {
        log.Fatal("Config file is missing: ", configfile)
    }
    var config Configuration
	  if _, err := toml.DecodeFile(configfile, &config); err != nil {
		    log.Fatal(err)
	  }
		return config
}
