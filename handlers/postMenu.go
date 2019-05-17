package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "database/sql"
    "time"
    "github.com/lib/pq"
    "os"
)

func PostMenu(w http.ResponseWriter, r *http.Request){
    currentTime:= time.Now()
    menu:= Menu{}
    status:= Confirm{
        Status: "Error",
    }
    err := json.NewDecoder(r.Body).Decode(&menu)
    if err != nil{
        panic(err)
    }
    //fmt.Println("Got these :",menu.day,menu.reg,menu.bf1,menu.bf2,menu.bf1,menu.lun2,menu.din1,menu.din2)

    db := GetDB()
    db.Query("")
    status.Status = "OK"
    statusJson,err := json.Marshal(status)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
