package handlers

import(
    //"fmt"
    "encoding/json"
    "net/http"
    "strconv"
)

func GetItems(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    status := Confirm{
        Status: "error",
        Text: "",
    }
    var item,items string
    var cost,id int
    rows,_ := db.Query("SELECT * FROM items;")
    for rows.Next(){
        ezz := rows.Scan(&id,&item,&cost)
        if(ezz == nil){
            items = items + strconv.Itoa(id) + ":" + item + ":" + strconv.Itoa(cost) + ","
        }
    }
    itemJson,errm := json.Marshal(status)
    if errm!=nil{
        panic(errm)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(itemJson)
}
