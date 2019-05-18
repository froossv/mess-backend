package handlers

import (
  "net/http"
  "encoding/json"
  "fmt"
  //"time"
  "github.com/lib/pq"
  "os"
  "database/sql"
)

func GetMenu(w http.ResponseWriter, r *http.Request){
    fmt.Println("We are in getMenu!")
    //currentDate := time.Now()
    //beforeDate := time.Now().AddDate(0,0,-1);
    //todayMenu := Menu{}
    //beforeMenu := Menu{}
    status := Confirm{
        Status: "Online",
        Text: "All right!",
    }
    dbUrl := os.Getenv("DATABASE_URL")
    psqlInfo,_ := pq.ParseURL(dbUrl)
    psqlInfo += " sslmode=require"
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
    //tquery := "SELECT * FROM menu WHERE day = '" + currentDate.Format("2006-01-02") + "';"
    //err1 := db.QueryRow(tquery).Scan(&todayMenu.day,&todayMenu.bf1,&todayMenu.bf1c,&todayMenu.bf2,&todayMenu.bf2c,&todayMenu.lun1,&todayMenu.lun1c,&todayMenu.lun2,&todayMenu.lun2c,&todayMenu.din1,&todayMenu.din1c,&todayMenu.din2,&todayMenu.din2c)
    //yquery := "SELECT * FROM menu WHERE day = '" + beforeDate.Format("2006-01-02") + "';"
    //err2 := db.QueryRow(yquery).Scan(&beforeMenu.day,&beforeMenu.bf1,&beforeMenu.bf1c,&beforeMenu.bf2,&beforeMenu.bf2c,&beforeMenu.lun1,&beforeMenu.lun1c,&beforeMenu.lun2,&beforeMenu.lun2c,&beforeMenu.din1,&beforeMenu.din1c,&beforeMenu.din2,&beforeMenu.din2c)
    //if (err1!=nil || err2!=nil){
    //    status.Text = "Not Good"
    //    panic(err1)
    //}else{
    //    fmt.Println(todayMenu)
    //    fmt.Println(beforeMenu)
    //}

    statusJson,err := json.Marshal(status)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
