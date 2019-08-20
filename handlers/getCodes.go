package handlers

import(
    "net/http"
    "encoding/json"
)

func GetCodes(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    codes := make([]CodeDet,0)
    rows,_ := db.Query("SELECT * FROM codes;")
    defer rows.Close()
    for rows.Next(){
        code := CodeDet{}
        _ = rows.Scan(&code.Username,&code.Code)
        codes = append(codes,code)
    }
    dataJson,errj := json.Marshal(codes)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(dataJson)
}
