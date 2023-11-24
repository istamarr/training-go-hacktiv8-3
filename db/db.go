package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 2345
	user   = "postgres"
	pwd    = "admin"
	dbname = "mar-db-go"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s pwd=%s 
	dbname=%s sslmode=disable`, host, port, user, pwd, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
