package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

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
