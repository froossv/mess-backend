package routing

import (
    "net/http"
    "mess-backend/handlers"
)

type Route struct{
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "Menu",
        "GET",
        "/menu",
        handlers.Menu,
    },
    Route{
        "Orders",
        "POST",
        "/orders",
        handlers.Orders,
    },
    Route{
        "Users",
        "POST",
        "/users",
        handlers.Users,
    },
}
