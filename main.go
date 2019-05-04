package main

import (
    "log"
    "fmt"
    "os"
    "net/http"
    "backendSastraMess/routing"
    //"github.com/gorilla/handlers"
    "github.com/rs/cors"
)

func getListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return port, nil
}

func main(){

    port,err := getListenAddress()
    fmt.Println(port)
    if err!= nil{
        log.Fatal(err)
    }
    router := routing.NewRouter()

    c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowCredentials: true,
    })
    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":"+port,handler))
    //log.Fatal(http.ListenAndServe(":"+port,handlers.CORS()(router)))
}
