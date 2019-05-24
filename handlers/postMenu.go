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
        Text: "Inserted Record at " + currentTime.String(),
    }
    err := json.NewDecoder(r.Body).Decode(&menu)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",menu.Bf1,menu.Bf1c,menu.Bf2,menu.Bf2c,menu.Lun1,menu.Lun1c,menu.Lun2,menu.Lun2c,menu.Din1,menu.Din1c,menu.Din2,menu.Din2c)
    menu.Day = currentTime.Format("2006-01-02")
    db := GetDB()
    defer db.Close()
    _,errr := db.Exec("INSERT INTO menu VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);",menu.Day,menu.Bf1,menu.Bf1c,menu.Bf2,menu.Bf2c,menu.Lun1,menu.Lun1c,menu.Lun2,menu.Lun2c,menu.Din1,menu.Din1c,menu.Din2,menu.Din2c)
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
