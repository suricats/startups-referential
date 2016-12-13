package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "io"
    "github.com/gorilla/mux"
    "github.com/suricats/startups-referential/models"
)

// IndexHandler handles requests to root api path (/)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Welcome, I am here to serve features around startups")
}

// GetStartupsHandler handles GET requests to /startups
func GetStartupsHandler(w http.ResponseWriter, r *http.Request) {
    startups := models.Startups{
        models.Startup{Name: "Dataiku"},
        models.Startup{Name: "Dreamquark"},
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(startups); err != nil {
        panic(err)
    }
}

// GetStartupHandler handles GET requests to /startups/{id}
func GetStartupHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    startupID := vars["startupId"]
    startup := models.Startup{Id: startupID, Name: "Dataiku"}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(startup); err != nil {
        panic(err)
    }
}

// CreateStartupHandler handles POST requests to /startups
func CreateStartupHandler(w http.ResponseWriter, r *http.Request) {
    var startup models.Startup
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &startup); err != nil {
       w.WriteHeader(http.StatusUnprocessableEntity)
       if err := json.NewEncoder(w).Encode(err); err != nil {
           panic(err)
       }
   }
   //should insert it into database and return created object
   s := startup
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusCreated)
   if err := json.NewEncoder(w).Encode(s); err != nil {
        panic(err)
   }
}
