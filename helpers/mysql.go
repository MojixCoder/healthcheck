package helpers

import (
	"database/sql"
	"fmt"
)

// MySQLConnResult is MySQL connection result
type MySQLConnResult struct {
	Conn       *sql.DB
	Error      error
	ConnString string
}

// ConnectToMySQL is for connecting to MySQL
func ConnectToMySQL(username, password, host, port, DBName string, ch chan MySQLConnResult) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username, password, host, port, DBName,
	)
	conn, err := sql.Open("mysql", connStr)
	ch <- MySQLConnResult{
		Conn:       conn,
		Error:      err,
		ConnString: connStr,
	}
}
