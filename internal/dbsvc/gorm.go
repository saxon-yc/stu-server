package dbsvc

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewGormContext(host string, port int, user, password, dbname string) *gorm.DB {
	gormConfig := gorm.Config{}
	gormConfig.NamingStrategy = schema.NamingStrategy{
		SingularTable: true,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		s := fmt.Sprintf("can't connect to database:%s,error:%s\n", dsn, err)
		log.Fatal(s)
	}
	sqlDb, err := db.DB()
	if err != nil {
		s := fmt.Sprintf("get connection pool failed, error: %s\n", err)
		log.Fatal(s)
	}
	sqlDb.SetConnMaxLifetime(8 * time.Second)
	sqlDb.SetMaxOpenConns(20)

	return db
}
