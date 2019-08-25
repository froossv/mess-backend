package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

//0 -> User
//1 -> Mess
//2 -> Nickname
func PostUser(w http.ResponseWriter, r *http.Request){
    cUser := UserDet{}
    muser := messUser{}
    suser := UserDet{}
    status:= Confirm{
        Status: "false",
        Text: "",
    }
    db := GetDB()
    defer db.Close()
    option := r.URL.Query()["user"][0]
    switch option {
        //user
        case "0": {
            err := json.NewDecoder(r.Body).Decode(&suser)
            if(err != nil){
                fmt.Println(err)
            }
            fmt.Println("Got these :",suser.Username,suser.Password,suser.Nickname)
            errr := db.QueryRow("SELECT reg,pwd,name,hostel,verified FROM users WHERE reg = $1;", suser.Username).Scan(&cUser.Username,&cUser.Password,&cUser.Nickname,&cUser.Hostel,&cUser.Verified)
            if errr!=nil{
                fmt.Println(errr)
                status.Status = "false"
            }else{
                if suser.Password == cUser.Password{
                    fmt.Println("Exists")
                    status.Status = "true"
                    status.Text = cUser.Nickname + "," + cUser.Hostel + "," + cUser.Verified
                }else{
                    fmt.Println("Wrong Password")
                    status.Status = "false"
                }
            }
        }
        //mess
        case "1": {
            err := json.NewDecoder(r.Body).Decode(&muser)
            if(err != nil){
                fmt.Println(err)
            }
            fmt.Println("Got these :",muser.Username,muser.Password)
            query := "SELECT username,password,name FROM mess WHERE username = '" + muser.Username + "';"
            fmt.Println(query)
            errs := db.QueryRow(query).Scan(&cUser.Username,&cUser.Password,&cUser.Nickname)
            if errs!=nil{
                fmt.Println(errs)
                status.Status = "false"
            }else{
                if muser.Password == cUser.Password{
                    fmt.Println("Exists")
                    status.Status = "true"
                    status.Text = cUser.Nickname
                }else{
                    fmt.Println("Wrong Password")
                    status.Status = "false"
                }
            }
        }
        //Nickname
        case "2": {
            err := json.NewDecoder(r.Body).Decode(&suser)
            if(err != nil){
                fmt.Println(err)
            }
            fmt.Println("Got these :",suser.Username,suser.Password,suser.Nickname)
            _,errp := db.Exec("UPDATE users SET name = $1 WHERE reg = $2;",suser.Nickname,suser.Username)
            if errp == nil{
                status.Status = "true"
                status.Text = suser.Nickname
            }else{
                fmt.Println(errp)
                status.Status = "false"
            }
        }
    }
    statusJson,err := json.Marshal(status)
    if(err != nil){
        fmt.Println(err)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
