package handlers

import (
    "encoding/json"
    "net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request){
    db := GetDB();
    status := Confirm{}
    reg := r.URL.Query()["reg"][0]
    var creg int
    err := db.QueryRow("SELECT reg FROM pwi WHERE reg = $1;",reg).Scan(&creg)
    if(err != nil){
        //no pwi
        status.Status = "false"
        statusJson,err := json.Marshal(status)
        if err!=nil{
            panic(err)
        }
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(statusJson)
    }else{
        //yes pwi
        status.Status = "true"
        errr := db.QueryRow("SELECT reg FROM users WHERE reg = $1;",reg).Scan(&creg)
        if(errr != nil){
            //no users
            status.Text = "false"
        }else{
            //yes users
            status.Text = "true"
        }
        statusJson,err := json.Marshal(status)
        if err!=nil{
            panic(err)
        }
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(statusJson)
    }
}
