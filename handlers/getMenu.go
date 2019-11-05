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
    var tm,to Menu
    menu := SendMenu{
        TBf: "null",
        TLun: "null",
        TDin: "null",
        TSnk: "null",
        NBf: "null",
        NLun: "null",
        NDin: "null",
        NSnk: "null",
    }
    date := time.Now()
    reg := r.URL.Query()["reg"][0]
    fmt.Println(reg)
    fmt.Println(date.Format("2006-01-02"))
    querym := "SELECT bf,lun,din,snk FROM menu WHERE day = '" + date.Format("2006-01-02") + "';"
    errm := db.QueryRow(querym).Scan(&tm.Bf,&tm.Lun,&tm.Din,&tm.Snk)
    queryo := "SELECT bf,lun,din,snk FROM orders WHERE day = '" + date.Format("2006-01-02") + "' AND reg = " + reg + ";"
    erro := db.QueryRow(queryo).Scan(&to.Bf,&to.Lun,&to.Din,&to.Snk)

    if(errm != nil || erro != nil){
        fmt.Println(errm)
        fmt.Println(erro)
    }

    fmt.Println(tm.Lun);
    fmt.Println(strings.Split(tm.Lun,","))
    fmt.Println(len(strings.Split(tm.Lun,",")))

    // for x,i := range len(strings.Split(tm.Bf,",")) {
    //
    // }

    menuJson,err := json.Marshal(menu)
    if(err!=nil){
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(menuJson)
}
