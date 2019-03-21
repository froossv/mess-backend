package main

import (
    "log"
    "fmt"
    "os"
    "net/http"
    "backendSastraMess/routing"
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
    log.Fatal(http.ListenAndServe(":"+port,router))
}
