package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const dsn = "host=rc1a-zyfnukz4le4di0xe.mdb.yandexcloud.net " +
	"port=6432 " +
	"user=aboba " +
	"password=123456789z " +
	"dbname=etl " +
	"sslmode=verify-full " +
	"sslrootcert=C:/Users/Sergc.DESKTOP-TU3FSM6/.postgresql/root.crt"

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil, err
	}

	log.Println("Connected to DB!")
	return db, nil

}
