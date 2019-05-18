package routing

import (
    "net/http"
    "backendSastraMess/handlers"
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
        "getOrders",
        "GET",
        "/orders",
        handlers.GetOrders,
    },
    Route{
        "postOrders",
        "POST",
        "/orders",
        handlers.PostOrders,
    },
    Route{
        "getMenu",
        "GET",
        "/menu",
        handlers.GetMenu,
    },
    Route{
        "postMenu",
        "POST",
        "/menu",
        handlers.PostMenu,
    },
    Route{
        "Users",
        "POST",
        "/users",
        handlers.PostUser,
    },
    Route{
        "Users",
        "PUT",
        "/users",
        handlers.PutUser,
    },
}
