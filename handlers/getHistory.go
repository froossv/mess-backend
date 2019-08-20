package handlers

import(
    //"fmt"
    //"encoding/json"
    "net/http"
)

func GetHistory(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()

}
