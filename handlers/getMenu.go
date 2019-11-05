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
    if(to.Bf != "null"){
        for i,x := range strings.Split(tm.Bf,",") {
            Bfos := strings.Split(to.Bf,",")
            if(Bfos[i] != "null"){
                menu.TBf = menu.TBf + x + "-" + GenRand(5) + Bfos[i] + ","
            }else{
                menu.TBf = menu.TBf + x + "-" + "null" + ","
            }
        }
    }
    fmt.Println(strings.TrimSuffix(menu.TBf,","))

    menuJson,err := json.Marshal(menu)
    if(err!=nil){
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(menuJson)
}
