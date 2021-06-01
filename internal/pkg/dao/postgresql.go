package dao

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open() *gorm.DB {

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")
	sslmode := os.Getenv("PG_SSLMODE")

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode + " Timezone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
