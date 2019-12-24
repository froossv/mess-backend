package handlers

import(
    //"fmt"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

func GetItems(w http.ResponseWriter, r *http.Request){
    db := GetDB()
    defer db.Close()
    status := Confirm{
        Status: "error",
        Text: "",
    }
    var item,items,deleted string
    var cost,id int
    rows,err := db.Query("SELECT * FROM items WHERE deleted = 'false' ORDER BY id ASC;")
    if(err == nil){
        for rows.Next(){
            ezz := rows.Scan(&id,&item,&cost,&deleted)
            if(ezz == nil){
                items = items + strconv.Itoa(id) + ":" + item + ":" + strconv.Itoa(cost) + ":" + deleted + ","
            }
        }
        status.Status = "Menu Items"
        status.Text = strings.TrimSuffix(items,",")
    }
    itemJson,errm := json.Marshal(status)
    if errm!=nil{
        panic(errm)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(itemJson)
}
