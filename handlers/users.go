package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/lib/pq"
    "os"
    "database/sql"
)

//type Confirm struct {
//    Status string `json: status`
//}

type UserDet struct{
    Username int `json: username`
    Password int `json: password`
}

func enableCors(w *http.ResponseWriter, r *http.Request){
    (*w).Header().Set("Access-Control-Allow-Origin","*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Users(w http.ResponseWriter, r *http.Request){
    //enableCors(&w,r);
    cUser:=UserDet{}
    user:=UserDet{}
    status:= Confirm{
        Status: "false",
    }

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",user.Username,user.Password)
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

    errr := db.QueryRow("SELECT reg,pwd FROM users WHERE reg = $1;", user.Username).Scan(&cUser.Username,&cUser.Password)
    if errr!=nil{
    }else{
        if user.Password == cUser.Password{
            fmt.Println("Exists")
            status.Status = "true"
        }else{
            fmt.Println("Doesnt Exist")
            status.Status = "false"
        }
    }

    statusJson,err := json.Marshal(status)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
