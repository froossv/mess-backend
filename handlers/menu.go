package handlers

import (
    "fmt"
    "net/http"
    //"encoding/json"
)

func Menu(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Menu is Up!")
}
