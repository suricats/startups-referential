package config

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
    Database string
}

type Configuration struct {
    Database DatabaseConfig `toml:"database"`
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
