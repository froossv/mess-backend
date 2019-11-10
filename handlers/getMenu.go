package handlers

import (
  "net/http"
  "encoding/json"
  "time"
  "fmt"
  "strings"
)

//date = time.Now().AddDate(0,0,1)
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
    menu := SendMenu{
        TBf: "",
        TLun: "",
        TDin: "",
        TSnk: "",
        NBf: "",
        NLun: "",
        NDin: "",
        NSnk: "",
    }
    date := time.Now()
    reg := r.URL.Query()["reg"][0]
    // fmt.Println(reg)
    fmt.Println(date.Format("2006-01-02"))

    querym := "SELECT bf,lun,din,snk FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errm := db.QueryRow(querym).Scan(&tm.Bf,&tm.Lun,&tm.Din,&tm.Snk)
    queryo := "SELECT bf,lun,din,snk FROM orders WHERE day = '" + date.Format("2006-01-02") + "' AND reg = " + reg + ";"
    erro := db.QueryRow(queryo).Scan(&to.Bf,&to.Lun,&to.Din,&to.Snk)
    fmt.Println(tm)
    fmt.Println(to)
    if(errm != nil || erro != nil){
        fmt.Println(errm)
        fmt.Println(erro)
    }
    writeToJson(&menu.TBf,tm.Bf,to.Bf)
    writeToJson(&menu.TLun,tm.Lun,to.Lun)
    writeToJson(&menu.TDin,tm.Din,to.Din)
    writeToJson(&menu.TSnk,tm.Snk,to.Snk)

    date = time.Now().AddDate(0,0,1)
    querym = "SELECT bf,lun,din,snk FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errm = db.QueryRow(querym).Scan(&tm.Bf,&tm.Lun,&tm.Din,&tm.Snk)
    queryo = "SELECT bf,lun,din,snk FROM orders WHERE day = '" + date.Format("2006-01-02") + "' AND reg = " + reg + ";"
    erro = db.QueryRow(queryo).Scan(&to.Bf,&to.Lun,&to.Din,&to.Snk)
    fmt.Println(tm)
    fmt.Println(to)
    if(errm != nil || erro != nil){
        fmt.Println(errm)
        fmt.Println(erro)
    }
    writeToJson(&menu.NBf,tm.Bf,to.Bf)
    writeToJson(&menu.NLun,tm.Lun,to.Lun)
    writeToJson(&menu.NDin,tm.Din,to.Din)
    writeToJson(&menu.NSnk,tm.Snk,to.Snk)

    menuJson,err := json.Marshal(menu)
    if(err!=nil){
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(menuJson)
}

func writeToJson(Dest *string,menu,order string) string{
    if(menu != "null"){
        for i,x := range strings.Split(menu,","){
            if(order != "null"){
                Bfos := strings.Split(order,",")
                if(Bfos[i] != "null"){
                    *Dest = *Dest + x + "-" + GenRand(5) + Bfos[i] + ","
                }else{
                    *Dest = *Dest + x + "-" + "null" + ","
                }
            }else{
                *Dest = *Dest + x + "-" + "null" + ","
            }
        }
    }else{
        *Dest = "null"
    }
    return strings.TrimSuffix(*Dest,",")
}
