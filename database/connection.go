package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Konek ke database
func Connect() *sql.DB {
	// Memiliki 2 parameter mandatory, nama driver dan connection string.
	var db, err = sql.Open(
		"mysql",
		"root@tcp(localhost:3306)/sekolah_db",
	)
	// Error Handling
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}
