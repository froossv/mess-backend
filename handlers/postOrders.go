package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
    "reflect"
)

func PostOrders(w http.ResponseWriter, r *http.Request){
    date := time.Now()
    order := Order{}
    status := Confirm{
        Status: "error",
        Text: "",
    }
    var costs [6]int64
    var cost int64
    err := json.NewDecoder(r.Body).Decode(&order)
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these:",order)

    db:= GetDB();

    errw := db.QueryRow("SELECT bf1c,bf2c,lun1c,lun2c,din1c,din2c FROM menu WHERE day = $1;",date.Format("2006-01-02")).Scan(&costs[0],&costs[1],&costs[2],&costs[3],&costs[4],&costs[5])
    if errw != nil{
        fmt.Println(errw)
    }
    val := reflect.ValueOf(order)
    for i := 1; i < val.NumField()-1; i++ {
        cost = cost + val.Field(i).Int() * costs[i-1]
    }
    query := "INSERT INTO orders VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) ON CONFLICT (regno) DO UPDATE set bf1 = excluded.bf1,bf2 = excluded.bf2,lun1 = excluded.lun1, lun2 = excluded.lun2, din1 = excluded.din1, din2 = excluded.din2;"
    _,erre := db.Exec(query,date.AddDate(0,0,1).Format("2006-01-02"),order.Username,order.Bf1,order.Bf2,order.Lun1,order.Lun2,order.Din1,order.Din2,cost)
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
