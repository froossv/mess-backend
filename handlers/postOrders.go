package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
)

func PostOrders(w http.ResponseWriter, r *http.Request){
    date := time.Now()
    order := Order{}
    status := Confirm{
        Status: "error",
        Text: "",
    }
    err := json.NewDecoder(r.Body).Decode(&order)
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these:",order)

    db:= GetDB();
    tableName := CreateTable()
    fmt.Println(tableName)
    _,erre := db.Exec("INSERT INTO "+tableName+" VALUES ($1,$2,$3,$4,$5,$6,$7,$8);",date.Format("2006-01-02"),order.Username,order.Bf1,order.Bf2,order.Lun1,order.Lun2,order.Din1,order.Din2)
    if(erre == nil){
        status.Status = "true"
    }else{
        fmt.Println(erre)
    }
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
