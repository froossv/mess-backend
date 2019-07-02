package handlers

import (
    "net/http"
    "encoding/json"
    "time"
)

func Index(w http.ResponseWriter, r *http.Request){
    cTime := time.Now()
    status := Confirm{
       Status: "Online",
       Text: cTime.Format("2006-01-02 15:04:05"),
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    if serr := json.NewEncoder(w).Encode(status); serr != nil{
        panic(serr)
    }
}
