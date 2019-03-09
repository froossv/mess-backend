package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "database/sql"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

type Confirm struct {
    Status string `json: status`
}

type studOrder struct{
    Reg int `json:reg`
    Bf string `json:bf`
    Lun string `json:lun`
    Din string `json:din`
}

type UserDet struct{
    Username int `json: username`
    Password int `json: password`
}

func enableCors(w *http.ResponseWriter, r *http.Request){
    (*w).Header().Set("Access-Control-Allow-Origin","*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Index(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Server is Up")
    status := Confirm{
       Status: "Online",
    }
    w.Header().Set("Content-Type","application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if serr := json.NewEncoder(w).Encode(status); serr != nil{
        panic(serr)
    }
}

func Menu(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Menu is Up!")
}

func Orders(w http.ResponseWriter, r *http.Request){
    enableCors(&w,r)
    order:= studOrder{}
    status:= Confirm{
        Status: "Error",
    }

    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",order.Reg,order.Bf,order.Lun,order.Din)

    db,err := sql.Open("mysql","vathsan:mysqlrox@tcp(127.0.0.1:3306)/orders")
    if err! = nil{
        panic(err)
    }

    statusJson,err := json.Marshal(status)
    if err!=nil{
        panic(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}

func Users(w http.ResponseWriter, r *http.Request){
    enableCors(&w,r)
    user:= UserDet{}
    cUser:=UserDet{}
    status:= Confirm{
        Status: "false",
    }

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",user.Username,user.Password)

    db,err := sql.Open("mysql","vathsan:mysqlrox@tcp(127.0.0.1:3306)/users")
    if err!=nil{
        panic(err)
    }
    defer db.Close()
    errr := db.QueryRow("SELECT regno,pwd FROM students WHERE regno = ?", user.Username).Scan(&cUser.Username,&cUser.Password)
    if errr!=nil{
    }else{
        if user.Password == cUser.Password{
            fmt.Println("Exists")
            status.Status = "true"
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
