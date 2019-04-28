package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "database/sql"
    "time"
    "github.com/lib/pq"
    "os"
)

type studOrder struct{
    day string `json:day`
    reg int `json:reg`
    bf1 string `json:bf`
    bf2 string `json:bf`
    lun1 string `json:lun`
    lun2 string `json:lun`
    din1 string `json:din`
    din2 string `json:din`
}

func Orders(w http.ResponseWriter, r *http.Request){
    currentTime:= time.Now()
    order:= studOrder{}
    status:= Confirm{
        Status: "Error",
    }

    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",order.day,order.reg,order.bf1,order.bf2,order.bf1,order.lun2,order.din1,order.din2)

    dburl := os.Getenv("DATABASE_URL")
    psqlInfo,_ := pq.ParseURL(dburl)
    psqlInfo += " sslmode=require"
    //fmt.Println(psqlInfo)
    db,erro := sql.Open("postgres",psqlInfo)
    if erro != nil{
        fmt.Printf("Error in Validating")
        panic(erro)
    }
    defer db.Close()
    errp := db.Ping()
    if errp != nil{
        fmt.Printf("Error in Connecting")
        panic(errp)
    }

    fmt.Println(currentTime.Format("2006_01_02"))
    rows,errq := db.Query("SELECT 1 FROM 2019_03_20")
    if errq == nil{
        fmt.Println("No Error")
        fmt.Println(errq)
        defer rows.Close()
    }else{
        fmt.Println("Error")
        defer rows.Close()
    }

    status.Status = "OK"
    statusJson,err := json.Marshal(status)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
