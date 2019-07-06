package handlers

import (
  "net/http"
  "encoding/json"
  "time"
  "fmt"
)

//0 -> Today's Menu
//1 -> Tomorrow's Menu
//2 -> Return Menu List

func GetMenu(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    menu := Menu{
        Bf1: "null",
        Bf1c: 0,
        Bf2: "null",
        Bf2c: 0,
        Lun1: "null",
        Lun1c: 0,
        Lun2: "null",
        Lun2c: 0,
        Din1: "null",
        Din1c: 0,
        Din2: "null",
        Din2c: 0,
    }
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
            menu.Day = date.Format("2006-01-02")
        case "1":
            date = time.Now().AddDate(0,0,1)
            menu.Day = date.Format("2006-01-02")
        case "2":
            item := returnMenu()
            status.Status = "menu"
            status.Text = item
            itemJson,errm := json.Marshal(status)
            if errm!=nil{
                panic(errm)
            }
            w.Header().Set("Content-Type","application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(itemJson)
            return
    }
    query := "SELECT * FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errr := db.QueryRow(query).Scan(&menu.Day,&menu.Bf1,&menu.Bf1c,&menu.Bf2,&menu.Bf2c,&menu.Lun1,&menu.Lun1c,&menu.Lun2,&menu.Lun2c,&menu.Din1,&menu.Din1c,&menu.Din2,&menu.Din2c)
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

func returnMenu() string{
    fmt.Println("Return Menu")
    var item string
    var items string = " "
    db := GetDB()
    rows,erri := db.Query("SELECT items FROM items")
    if erri != nil{
        fmt.Println(erri)
    }
    for rows.Next(){
        ezz := rows.Scan(&item)
        if(ezz == nil){
            items = items + item + ","
        }
    }
    return items
}
