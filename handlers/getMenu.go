package handlers

import (
  "net/http"
  "encoding/json"
  "time"
  "fmt"
)

//0 -> Today's Menu
//1 -> Tomorrow's Menu

func GetMenu(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    menu := Menu{
        Bf: "null",
        Lun: "null",
        Din: "null",
        Snk: "null",
    }
    date := time.Now()
    option := r.URL.Query()["day"][0]
    fmt.Println(option)
    switch option {
        case "0":
            date = time.Now()
            menu.Day = date.Format("2006-01-02")
        case "1":
            date = time.Now().AddDate(0,0,1)
            menu.Day = date.Format("2006-01-02")
    }
    query := "SELECT * FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errr := db.QueryRow(query).Scan(&menu.Day,&menu.Bf,&menu.Lun,&menu.Din,&menu.Snk)
    if(errr != nil){
        fmt.Println(errr)
    }
    menuJson,err := json.Marshal(menu)
    if(err!=nil){
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(menuJson)
}
