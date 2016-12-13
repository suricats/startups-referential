package db

import (
    "github.com/suricats/startups-referential/config"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
)

func GetConnection(config config.Configuration) *sql.DB {
    db, err := sql.Open("mysql", config.Database.User + ":" + config.Database.Password + "@/" + config.Database.Database + "?charset=utf8")
    if err != nil {
        panic (err.Error())
    }
    err = db.Ping()
    if err != nil {
        panic (err.Error())
    }
    return db
}

func ListStartups(config config.Configuration, db *sql.DB) {
    var id int
    var name string
    stmt, err := db.Prepare("SELECT * FROM startup WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()
    rows, err := stmt.Query(1)
    if err != nil {
        panic(err.Error())
    }
    for rows.Next() {
      err := rows.Scan(&id, &name)
      if err != nil {
        log.Fatal(err)
      }
      log.Println(id, name)
    }
}
