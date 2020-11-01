package model

import (
	"log"

	"gorm.io/gorm"
)

// PopulateDB for Model
func (item Model) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for BusinessFunction
func (item BusinessFunction) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for SecurityPractice
func (item SecurityPractice) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for PracticeLevel
func (item PracticeLevel) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for MaturityLevel
func (item MaturityLevel) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for ActivityStream
func (item ActivityStream) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for Activity
func (item Activity) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}

// PopulateDB for Question
func (item Question) PopulateDB(db *gorm.DB) {
	//spew.Sdump(item)
	if result := db.Create(&item); result.Error != nil {
		log.Println(result.Error)
	}
}
