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
    //tableName := CreateTable()
    //fmt.Println(tableName)
    query := "INSERT INTO orders VALUES ($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT (regno) DO UPDATE set bf1 = excluded.bf1,bf2 = excluded.bf2,lun1 = excluded.lun1, lun2 = excluded.lun2, din1 = excluded.din1, din2 = excluded.din2;"
    _,erre := db.Exec(query,date.AddDate(0,0,1).Format("2006-01-02"),order.Username,order.Bf1,order.Bf2,order.Lun1,order.Lun2,order.Din1,order.Din2)
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
