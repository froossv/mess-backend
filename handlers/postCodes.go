package handlers

import(
    "fmt"
    "encoding/json"
    "net/http"
)

func VerUser(w http.ResponseWriter, r *http.Request){
    code := CodeDet{}
    var dbcode string
    status := Confirm{
        Status: "Error",
        Text: "",
    }
    err := json.NewDecoder(r.Body).Decode(&code);
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these :",code.Username,code.Code)
    db := GetDB();
    error := db.QueryRow("SELECT code FROM codes WHERE reg = $1",code.Username).Scan(&dbcode)
    if(error!=nil){
        if(code.Code == dbcode){
            status.Status = "true"
            _,erre := db.Exec("UPDATE users SET verified = 'true' WHERE reg = $1",code.Username)
            if(erre != nil){
                panic(erre)
            }
        }else{
            status.Status = "false"
        }
    }
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
