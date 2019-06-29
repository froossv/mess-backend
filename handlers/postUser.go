package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

//0 -> User
//1 -> Mess

func PostUser(w http.ResponseWriter, r *http.Request){
    cUser:=UserDet{}
    user:=UserDet{}
    status:= Confirm{
        Status: "false",
        Text: "",
    }
    db := GetDB()
    defer db.Close()
    option := r.URL.Query()["user"][0]
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil{
        panic(err)
    }
    fmt.Println("Got these :",user.Username,user.Password)
    switch option {
        case "0": {
            errr := db.QueryRow("SELECT reg,pwd,name,hostel,verified FROM users WHERE reg = $1;", user.Username).Scan(&cUser.Username,&cUser.Password,&cUser.Name,&cUser.Hostel,&cUser.Verified)
            if errr!=nil{
                fmt.Println(errr)
                status.Status = "false"
            }else{
                if user.Password == cUser.Password{
                    fmt.Println("Exists")
                    status.Status = "true"
                    status.Text = cUser.Name + "," + cUser.Hostel + "," + cUser.Verified
                }else{
                    fmt.Println("Wrong Password")
                    status.Status = "false"
                }
            }
        }
        case "1": {
            errs := db.QueryRow("SELECT username,password,name FROM mess WHERE username = $1;", user.Username).Scan(&cUser.Username,&cUser.Password,&cUser.Name)
            if errs!=nil{
                fmt.Println(errs)
                status.Status = "false"
            }else{
                if user.Password == cUser.Password{
                    fmt.Println("Exists")
                    status.Status = "true"
                    status.Text = cUser.Name
                }else{
                    fmt.Println("Wrong Password")
                    status.Status = "false"
                }
            }
        }

    }
    statusJson,err := json.Marshal(status)
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
