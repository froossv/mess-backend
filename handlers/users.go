package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

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
    fmt.Println("Got these :",user.Username,user.Password)
    db := GetDB();
    errr := db.QueryRow("SELECT reg,pwd,name,hostel FROM users WHERE reg = $1;", user.Username).Scan(&cUser.Username,&cUser.Password,&cUser.Name,&cUser.Hostel)
    if errr!=nil{
        fmt.Println(errr)
        panic(errr)
    }else{
        if user.Password == cUser.Password{
            fmt.Println("Exists")
            status.Status = "true"
            status.Text = cUser.Name + "," + cUser.Hostel
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
