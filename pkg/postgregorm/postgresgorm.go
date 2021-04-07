package postgresgorm

import (
	"fmt"
	"nuryanto2121/cukur_in_web/pkg/setting"

	// "job_cukur-in/pkg/setting"
	// util "job_cukur-in/pkg/utils"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func Setup() {
	now := time.Now()
	var err error

	connectionstring := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		setting.FileConfigSetting.Database.Host,
		setting.FileConfigSetting.Database.User,
		setting.FileConfigSetting.Database.Password,
		setting.FileConfigSetting.Database.Name,
		setting.FileConfigSetting.Database.Port)
	fmt.Printf("%s", connectionstring)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Jakarta"
	Conn, err = gorm.Open(postgres.Open(connectionstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.FileConfigSetting.Database.TablePrefix,
			SingularTable: true,
		},
		PrepareStmt: true,
		Logger:      newLogger,
		// DryRun: true,
	})

	if err != nil {
		log.Printf("connection.setup err : %v", err)
		panic(err)
	}

	sqlDB, err := Conn.DB()
	if err != nil {
		log.Printf("connection.setup DB err : %v", err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// sqlDB.Sing
	// Conn.DB().SetMaxIdleConns(10)
	// Conn.DB().SetMaxOpenConns(100)

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
