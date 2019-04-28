package routing

import(
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)


func NewRouter() *mux.Router{
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
