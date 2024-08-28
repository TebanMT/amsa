package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func GetDatabaseInstance() *gorm.DB {
	once.Do(func() {
		dsn := "admin:Admin2024*@tcp(terraform-20240107183529051700000001.cn6om8com5b4.us-east-1.rds.amazonaws.com:3306)/amsa?parseTime=true" //os.Getenv("DATABASE_DSN")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("No se pudo conectar a la base de datos: %v", err)
		}
		instance = db
	})
	return instance
}
