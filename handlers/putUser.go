package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func PutUser(w http.ResponseWriter, r *http.Request){
    user := UserDet{}
    pUser := UserDet{}
    status := Confirm{
        Status: "Error",
        Text: "",
    }
    err := json.NewDecoder(r.Body).Decode(&user);
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these :",user.Username,user.Password,user.Nickname)
    db := GetDB()
    defer db.Close()

    //check if already exists in users
    errr := db.QueryRow("SELECT reg FROM users WHERE reg = $1",user.Username)
    if(errr != nil){
        //he no exist
        fmt.Println("!Users")
        //check if he exists in pwi
        errp := db.QueryRow("SELECT reg,hostel FROM pwi WHERE reg = $1",user.Username).Scan(&pUser.Username,&pUser.Hostel)
        if(errp != nil){
            //he no exists in pwi
            fmt.Println("!PWI")
            goto EXIT
        }else{
            //he exist in pwi
            fmt.Println("PWI")
            //insert him into ours
            _,erra := db.Exec("INSERT INTO users VALUES ($1,$2,$3,$4,'false');",user.Username,user.Password,user.Nickname,pUser.Hostel)
            code := GenRand(8)
            _,errc := db.Exec("INSERT INTO codes VALUES ($1,$2);",user.Username,code)
            if(erra!=nil && errc!=nil){
                //insert him to users failed
                goto EXIT
            }else{
                fmt.Println("Users")
                //inserted him into users
                status.Status = "OK"
                goto EXIT
            }
        }
    }else if(errr == nil){
        //he exist
        fmt.Println("Error is not nil")
        goto EXIT
    }
    EXIT:
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
