package handlers

import(
    "github.com/lib/pq"
    "database/sql"
    "os"
    "fmt"
    "time"
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

func CreateTable() string{
    date := time.Now()
    db := GetDB()
    defer db.Close()
    query := "CREATE TABLE IF NOT EXISTS orders_" + date.Month().String() + "(day date,reg int,bf1 int,bf2 int,lun1 int,lun2 int,din1 int,din2 int)"
    _,err := db.Exec(query)
    if(err!=nil){
        return "error"
    }
    return "orders_" + date.Month().String()
}

func G89-enRand(n int) string {
    var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, n)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}
