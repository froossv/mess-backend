package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)
type Dummy struct {
    status string
    text string
}

func Index(w http.ResponseWriter, r *http.Request){
    test := Dummy{

    }
}

func Menu(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Menu is Up!")
}
func Orders(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Orders is Up!")
}
