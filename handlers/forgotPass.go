package handlers

import(
    "encoding/json"
    "net/http"
)

func ForgotPass(w http.ResponseWriter, r *http.Request){
    passwd := pass{}
    db := GetDB()
    status := Confirm{
        Status: "true",
        Text: "",
    }

    err := json.NewDecoder(r.Body).Decode(&passwd)
    if err!= nil{
        panic(err)
    }
    _,erra := db.Exec("UPDATE users SET pwd = $1 WHERE reg = $2",passwd.New,passwd.Username)
    if erra!=nil{
        status.Status = "false"
    }
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
