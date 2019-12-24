package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
    "strconv"
    "reflect"
    "strings"
)

func PostOrders(w http.ResponseWriter, r *http.Request){
    date := time.Now()
    order := Order{}
    status := Confirm{
        Status: "error",
        Text: "",
    }
    var cost int
    err := json.NewDecoder(r.Body).Decode(&order)
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these:",order)
    db:= GetDB();
    val := reflect.ValueOf(order)
    for i := 1; i < val.NumField(); i++ {
        if(val.Field(i).String() != "null"){
            for _,x := range strings.Split(val.Field(i).String(),","){
                if(x != "null"){
                    iqa := strings.Split(x,"-")
                    q,_ := strconv.Atoi(iqa[1])
                    a,_ := strconv.Atoi(iqa[2])
                    cost = cost + q*a
                }
            }
        }
    }
    query := "INSERT INTO orders VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT ON CONSTRAINT \"uniqueregdate\" DO UPDATE set bf = excluded.bf,lun = excluded.lun,din = excluded.din, snk = excluded.snk;"
    _,erre := db.Exec(query,date.Format("2006-01-02"),order.Username,order.Bf,order.Lun,order.Din,order.Snk,cost)
    if(erre == nil){
         status.Status = "true"
    }else{
         fmt.Println(erre)
    }
    var codes [4]string
    for i := 1; i < val.NumField(); i++ {
        if(val.Field(i).String() != "null"){
            for _,x := range strings.Split(val.Field(i).String(),","){
                if(x != "null"){
                    iqa := strings.Split(x,"-")
                    id,_ := strconv.Atoi(iqa[0])
                    q,_ := strconv.Atoi(iqa[1])
                    codes[i-1] += "," + GenRand(5) + fmt.Sprintf("%02d",id) + fmt.Sprintf("%02d",q)
                }
            }
        }else{
            codes[i-1] = "null"
        }
        codes[i-1] = strings.Trim(codes[i-1],",")
    }
    _,errt := db.Exec("INSERT INTO order_codes values($1,$2,$3,$4,$5) ON CONFLICT (reg) DO UPDATE set bf = excluded.bf,lun = excluded.lun,din = excluded.din,snk = excluded.snk;",order.Username,codes[0],codes[1],codes[2],codes[3])
    if(errt == nil){
        status.Status = "true"
    }else{
        fmt.Println(errt)
    }
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
