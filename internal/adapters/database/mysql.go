package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/RiteshDevOpsEngineer/ecom/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlOnce sync.Once
	db        *Database
)

type Database struct {
	Conn *sql.DB
}

func InitializeDatabase() (*Database, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	// fmt.Println("Connected to the database")
	return &Database{Conn: conn}, nil
}

func GetDatabase() *Database {
	mysqlOnce.Do(func() {
		var err error
		db, err = InitializeDatabase()
		if err != nil {
			panic(err)
		}
	})
	return db
}
