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
        Bf1: "null",
        Bf2: "null",
        Lun1: "null",
        Lun2: "null",
        Din1: "null",
        Din2: "null"
    }
    var table string
    option := r.URL.Query()["day"][0]
    switch option {
        case "0":{
            table = "order_codes_today"
        }
        case "1":{
            table = "order_codes_tomorrow"
        }
    }
    query := "SELECT bf1,bf2,lun1,lun2,din1,din2 FROM " + table + " WHERE regno = " + r.URL.Query()["regno"][0] +";"
    erro := db.QueryRow(query).Scan(&codes.Bf1,&codes.Bf2,&codes.Lun1,&codes.Lun2,&codes.Din1,&codes.Din2)
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
