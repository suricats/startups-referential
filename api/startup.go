package main

import "time"

type Startup struct {
    Id          string    `json:"id"`
    Name        string    `json:"name"`
    LastContact time.Time `json:"last_contact"`
}

type Startups []Startup
