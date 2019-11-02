package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

//0 -> today
//0 -> tomorrow
func GetOrders(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    codes := Code{
        Bf: "null",
        Lun: "null",
        Din: "null",
        Snk: "null",
    }
    var name string
    option := r.URL.Query()["day"][0]
    switch option {
        case "0":{
            name = "order_codes_today"
        }
        case "1":{
            name = "order_codes_tomorrow"
        }
    }
    query := "SELECT bf,lun,din,snk FROM "+ name +" WHERE regno = " + r.URL.Query()["regno"][0] +";"
    fmt.Println(query)
    erro := db.QueryRow(query).Scan(&codes.Bf,&codes.Lun,&codes.Din,&codes.Snk)
    if erro != nil{
        fmt.Println(erro)
    }
    //codes.Username = regno
    codeJson,err := json.Marshal(codes)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(codeJson)

}
