package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/lib/pq"
    "os"
    "database/sql"
)

type UserDet struct{
    Username int `json: username`
    Password int `json: password`
    Hostel string `json: hostel`
}

func Users(w http.ResponseWriter, r *http.Request){
    cUser:=UserDet{}
    user:=UserDet{}
    status:= Confirm{
        Status: "false",
        Text: "",
    }

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",user.Username,user.Password,user.Hostel)
    dburl := os.Getenv("DATABASE_URL")
    psqlInfo,_ := pq.ParseURL(dburl)
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

    errr := db.QueryRow("SELECT reg,pwd,hostel FROM users WHERE reg = $1;", user.Username).Scan(&cUser.Username,&cUser.Password,&cUser.Hostel)
    if errr!=nil{
    }else{
        if user.Password == cUser.Password{
            fmt.Println("Exists")
            status.Status = "true"
            status.Text = cUser.Hostel
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
