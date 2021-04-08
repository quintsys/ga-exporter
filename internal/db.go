package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func conn() Connection, error {
	condb, errdb := sql.Open("mssql", "server=localhost;user id=sa;password=SA_PASSWORD=yourStrong(!)Password;")
	if errdb != nil {
		return nil, errdb.Error()
	}
}

func version(condb Connection) {
	var (
		sqlversion string
	)
	rows, err := condb.Query("select @@version")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sqlversion)
	}
	defer condb.Close()

}