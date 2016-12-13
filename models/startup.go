package models

type Startup struct {
    Id          string    `db:"id" json:"id"`
    Name        string    `db:"name" json:"name"`
}

type Startups []Startup
