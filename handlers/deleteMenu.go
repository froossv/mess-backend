package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"
)

func DeleteMenu(w http.ResponseWriter, r *http.Request){
    menu := CodeDet{}
    db := GetDB()
    status := Confirm{
        Status: "true",
        Text: "",
    }
    err := json.NewDecoder(r.Body).Decode(&menu)
    if(err!=nil){
        panic(err)
    }
    _,erra := db.Exec("DELETE FROM items WHERE items = $1;",menu.Code)
    if erra!=nil {
        status.Status  = "error"
        fmt.Println(erra)
    }
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
