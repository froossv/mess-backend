package handlers

import (
  "net/http"
  "encoding/json"
  "time"
  "fmt"
  "strings"
)

func GetMenu(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    tm := Menu{
        Bf: "null",
        Lun: "null",
        Din: "null",
        Snk: "null",
    }
    to := Menu{
        Bf: "null",
        Lun: "null",
        Din: "null",
        Snk: "null",
    }
    menu := Menu{
        Bf: "",
        Lun: "",
        Din: "",
        Snk: "",
    }
    reg := r.URL.Query()["reg"][0]
    fmt.Println(reg)
//date = time.Now().AddDate(0,0,1)
    date := time.Now()
    querym := "SELECT bf,lun,din,snk FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errm := db.QueryRow(querym).Scan(&tm.Bf,&tm.Lun,&tm.Din,&tm.Snk)
    queryo := "SELECT bf,lun,din,snk FROM order_codes WHERE reg = " + reg + ";"
    erro := db.QueryRow(queryo).Scan(&to.Bf,&to.Lun,&to.Din,&to.Snk)
    fmt.Println(tm)
    fmt.Println(to)
    if(errm != nil || erro != nil){
        fmt.Println(errm)
        fmt.Println(erro)
    }
    writeToJson(&menu.Bf,tm.Bf,to.Bf)
    writeToJson(&menu.Lun,tm.Lun,to.Lun)
    writeToJson(&menu.Din,tm.Din,to.Din)
    writeToJson(&menu.Snk,tm.Snk,to.Snk)

    menuJson,err := json.Marshal(menu)
    if(err!=nil){
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(menuJson)
}

func writeToJson(Dest *string,menu,order string){
    if(menu != "null"){
        for _,x := range strings.Split(menu,","){
            var flag bool
            *Dest += x
            if(order != "null"){
                for _,y := range strings.Split(order,","){
                    if(y[5:7] == x){
                        *Dest += "-" + y
                        flag = true
                        break
                    }
                }
                if(!flag){
                    *Dest += "-" + "null"
                }
            }else{
                *Dest += "-" + "null"
            }
            *Dest += ","
        }
    }else{
        *Dest = "null"
    }
    *Dest = strings.TrimSuffix(*Dest,",")
}
