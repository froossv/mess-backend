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
    query := "INSERT INTO menu VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) ON CONFLICT (day) DO UPDATE set bf1 = excluded.bf1, bf1c = excluded.bf1c, bf2 = excluded.bf2, bf2c = excluded.bf2c, lun1 = excluded.lun1,lun1c = excluded.lun1c, lun2 = excluded.lun2, lun2c = excluded.lun2c, din1 = excluded.din1, din1c = excluded.din1c, din2 = excluded.din2,din2c = excluded.din2c;"
    fmt.Println("Got these :",menu.Bf1,menu.Bf1c,menu.Bf2,menu.Bf2c,menu.Lun1,menu.Lun1c,menu.Lun2,menu.Lun2c,menu.Din1,menu.Din1c,menu.Din2,menu.Din2c)
    menu.Day = currentTime.Format("2006-01-02")
    db := GetDB()
    defer db.Close()
    _,errr := db.Exec(query,menu.Day,menu.Bf1,menu.Bf1c,menu.Bf2,menu.Bf2c,menu.Lun1,menu.Lun1c,menu.Lun2,menu.Lun2c,menu.Din1,menu.Din1c,menu.Din2,menu.Din2c)
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
