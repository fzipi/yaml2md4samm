package main

import (
	"log"
	"os"
	"time"
	"yaml2md4samm/model"
	"yaml2md4samm/parsing"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func initialMigration(DB *gorm.DB) {
	// Migrate the schema
	// DB.AutoMigrate(&model.Model{},
	// 	&model.BusinessFunction{},
	// 	&model.SecurityPractice{},
	// 	&model.PracticeLevel{},
	// 	&model.MaturityLevel{},
	// 	&model.ActivityStream{},
	// 	&model.Activity{},
	// 	&model.Question{})

	DB.AutoMigrate(&model.Model{})
	DB.AutoMigrate(&model.BusinessFunction{})
	DB.AutoMigrate(&model.SecurityPractice{})
	DB.AutoMigrate(&model.PracticeLevel{})
	DB.AutoMigrate(&model.MaturityLevel{})
	DB.AutoMigrate(&model.ActivityStream{})
	DB.AutoMigrate(&model.Activity{})
	DB.AutoMigrate(&model.Question{})
}

/* This is the relation between stuff in SAMM:

model > function > practice > stream > activity > question
*/
func main() {
	var samm = model.Model{Version: 2.0, License: "CC-1.0", ExecutiveSummary: "Summary"}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // Disable color
		},
	)

	DB, err := gorm.Open(sqlite.Open("samm.db"), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})

	if err != nil {
		panic("failed to create database samm.db")
	}

	logrus.Print("Performing initial DB migration.")
	initialMigration(DB)

	logrus.Print("Parsing SAMM model.")
	parsing.ParseFullModel(&samm, DB)
	logrus.Print("Parsed.")

}
