package handlers

import (
    "fmt"
    "net/http"
    //"encoding/json"
)

func GetOrders(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "postOrders is Up!")
}
