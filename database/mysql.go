package databse

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectManual() {
	connection, err := sql.Open("mysql", "root:root@/hotels?parseTime=true")

  if err != nil {
        log.Fatal(err)
	}

  db = connection
}



