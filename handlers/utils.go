package handlers

import(
    "github.com/lib/pq"
    "database/sql"
    "os"
    "fmt"
)

func GetDB() (*sql.DB){
    dburl := os.Getenv("DATABASE_URL")
    psqlInfo,_ := pq.ParseURL(dburl)
    psqlInfo += " sslmode=require"
    db,erro := sql.Open("postgres",psqlInfo)
    if erro != nil{
        fmt.Printf("Error in Validating")
        panic(erro)
    }
    errp := db.Ping()
    if errp != nil{
        fmt.Printf("Error in Connecting")
        panic(errp)
    }
    return db
}
