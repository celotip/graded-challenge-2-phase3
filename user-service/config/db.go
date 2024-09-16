package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	var (
		user     = "root"
		password = "tTlKPflgZSLPRcZpbLuQSeQTGFsMorIH"
		host     = "autorack.proxy.rlwy.net"
		port     = "11040"
		dbname   = "railway"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check connection with the database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret: "secret",
	}
}
