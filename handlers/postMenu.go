package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
)

func PostMenu(w http.ResponseWriter, r *http.Request){
    currentTime:= time.Now()
    menu:= Menu{}
    status:= Confirm{
        Status: "Error",
        Text: "Inserted Record on " + currentTime.String(),
    }
    err := json.NewDecoder(r.Body).Decode(&menu)
    if err != nil{
        panic(err)
    }
    query := "INSERT INTO menu VALUES ($1,$2,$3,$4,$5) ON CONFLICT (day) DO UPDATE set bf = excluded.bf, lun = excluded.lun,din = excluded.din,snk = excluded.snk;"
    fmt.Println("Got these :",menu.Bf,menu.Lun,menu.Din,menu.Snk)
    menu.Day = currentTime.Format("2006-01-02")
    fmt.Println(menu.Day)
    db := GetDB()
    defer db.Close()
    _,errr := db.Exec(query,menu.Day,menu.Bf,menu.Lun,menu.Din,menu.Snk)
    if errr!=nil{
        status.Text = "Error Inserting into table"
        panic(errr)
    }else{
        status.Status = "OK"
    }
    statusJson,err := json.Marshal(status)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
