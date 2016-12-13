package main

import (
    "net/http"
)

// Route defines a REST endpint
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes defines a REST routes set
type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        IndexHandler,
    },
    Route{
        "Get startups",
        "GET",
        "/startups",
        GetStartupsHandler,
    },
    Route{
        "Create startup",
        "POST",
        "/startups",
        CreateStartupHandler,
    },
    Route{
        "Get startup",
        "GET",
        "/startups/{startupId}",
        GetStartupHandler,
    },
}
