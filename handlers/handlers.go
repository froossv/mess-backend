package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Confirm struct {
    Status string `json: status`
}

func Index(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Server is Up")
    status := Confirm{
        Status: "Online",
    }
    w.Header().Set("Content-Type","application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if serr := json.NewEncoder(w).Encode(status); serr != nil{
        panic(serr)
    }
}

func Menu(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Menu is Up!")
}
func Orders(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Orders is Up!")
}
