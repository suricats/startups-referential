package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "io"
    "github.com/gorilla/mux"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Welcome, I am here to serve features around startups")
}

func GetStartupsHandler(w http.ResponseWriter, r *http.Request) {
    startups := Startups{
        Startup{Name: "Dataiku"},
        Startup{Name: "Dreamquark"},
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(startups); err != nil {
        panic(err)
    }
}

func GetStartupHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    startupId := vars["startupId"]
    startup := Startup{Id: startupId, Name: "Dataiku"}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(startup); err != nil {
        panic(err)
    }
}

func CreateStartupHandler(w http.ResponseWriter, r *http.Request) {
    var startup Startup
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &startup); err != nil {
       w.Header().Set("Content-Type", "application/json; charset=UTF-8")
       w.WriteHeader(422) // unprocessable entity
       if err := json.NewEncoder(w).Encode(err); err != nil {
           panic(err)
       }
   }
   s := startup
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusCreated)
   if err := json.NewEncoder(w).Encode(s); err != nil {
     panic(err)
   }
}
