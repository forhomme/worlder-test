package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	DbSql *gorm.DB
}

var (
	onceDbSql  sync.Once
	instanceDB *DB
)

func GetInstanceDb() *gorm.DB {
	onceDbSql.Do(func() {
		sqlInfo := Config.Database
		logs := fmt.Sprintf("[INFO] Connected to MYSQL TYPE = %s | LogMode = %+v", sqlInfo.Host, sqlInfo.LogMode)

		dbConfig := sqlInfo.Username + ":" + sqlInfo.Password + "@tcp(" + sqlInfo.Host + ":" + fmt.Sprintf("%d", sqlInfo.Port) +
			")/" + sqlInfo.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to MYSQL with err %s. Config=%s", err.Error(), sqlInfo.Host)
			log.Fatalln(logs)
		}

		sqlDB, err := db.DB()
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to MYSQL with err %s. Config=%s", err.Error(), sqlInfo.Host)
			log.Fatalln(logs)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetConnMaxLifetime(10 * time.Minute)
		dialect := mysql.New(mysql.Config{Conn: sqlDB})
		loggerLevel := logger.Error
		if sqlInfo.LogMode {
			loggerLevel = logger.Info
		}
		dbConnection, err := gorm.Open(dialect, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      loggerLevel,
				},
			),
		})
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to MYSQL with err %s. Config=%s", err.Error(), sqlInfo.Host)
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceDB = &DB{DbSql: dbConnection}
	})
	return instanceDB.DbSql
}
