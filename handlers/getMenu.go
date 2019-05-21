package handlers

import (
  "net/http"
  "encoding/json"
  "time"
  "fmt"
)

//0 -> Tomorrow's Menu
//1 -> Today's Menu

func GetMenu(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    menu := Menu{}
    status := Confirm{
        Status: "error",
        Text: "",
    }
    date := time.Now()

    option := r.URL.Query()["day"][0]
    fmt.Println(option)
    switch option {
        case "0":
            date = time.Now()
        case "1":
            date = time.Now().AddDate(0,0,1)
    }

    query := "SELECT * FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errr := db.QueryRow(query).Scan(&menu.Day,&menu.Bf1,&menu.Bf1c,&menu.Bf2,&menu.Bf2c,&menu.Lun1,&menu.Lun1c,&menu.Lun2,&menu.Lun2c,&menu.Din1,&menu.Din1c,&menu.Din2,&menu.Din2c)
    if(errr == nil){
        fmt.Println(menu)
        menuJson,err := json.Marshal(menu)
        if(err!=nil){
            panic(err)
        }
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(menuJson)
    }else{
        fmt.Println(errr)
        statusJson,err := json.Marshal(status)
        if err!=nil{
            panic(err)
        }
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(statusJson)
    }
}
