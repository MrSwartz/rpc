package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"rpc/db_server/internal/config"
)

func Connect(s *config.SQLConfig) *sql.DB {
	host := fmt.Sprintf("%s:%s@/%s", s.User, s.Password, s.Database)
	conn, err := sql.Open(s.Driver, host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Fprintf(os.Stdout, "%v : %v\n", s.Driver, host)
	return conn
}
