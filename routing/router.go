package routing

import(
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "net/http"
)


func NewRouter() http.Handler{
    router := mux.NewRouter()
    for _, route := range routes{
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    handler := cors.Default().Handler(router)
    return handler
}
