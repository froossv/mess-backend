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

func PostMenu(w http.ResponseWriter, r *http.Request){
    currentTime:= time.Now()
    menu:= Menu{}
    status:= Confirm{
        Status: "Error",
    }

    err := json.NewDecoder(r.Body).Decode(&menu)
    if err != nil{
        panic(err)
    }
    //fmt.Println("Got these :",menu.day,menu.reg,menu.bf1,menu.bf2,menu.bf1,menu.lun2,menu.din1,menu.din2)

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
