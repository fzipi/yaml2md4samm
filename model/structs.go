package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SAMMType is any of the model types
type SAMMType interface {
	PopulateDB(db *gorm.DB)
}

// https://yaml.to-go.online/

// Model is the base type
type Model struct {
	gorm.Model
	Version           float32 `yaml:"model" gorm:"primaryKey"`
	ExecutiveSummary  string
	License           string
	BusinessFunctions []BusinessFunction
}

// BusinessFunction Model has a list of
type BusinessFunction struct {
	ID                uuid.UUID `yaml:"id" gorm:"type:uuid;primaryKey"`
	ModelVersion      string    `yaml:"model" gorm:"foreignKey:ModelVersion"`
	Name              string    `yaml:"name"`
	Color             string    `yaml:"color"`
	Logo              string    `yaml:"logo"`
	Order             int8      `yaml:"order"`
	Description       string    `yaml:"description"`
	Type              string    `yaml:"type"`
	SecurityPractices []SecurityPractice
}

// SecurityPractice each Business Functions have a list of
type SecurityPractice struct {
	ID                 uuid.UUID `yaml:"id" gorm:"type:uuid;primaryKey"`
	Name               string    `yaml:"name"`
	ShortName          string    `yaml:"shortName"`
	BusinessFunctionID uuid.UUID `yaml:"function" gorm:"type:uuid"`
	Progress           string    `yaml:"progress"`
	Assignee           string    `yaml:"assignee"`
	Order              int8      `yaml:"order"`
	ShortDescription   string    `yaml:"shortDescription"`
	LongDescription    string    `yaml:"longDescription" gorm:"type:text"`
	Type               string    `yaml:"type"`
	PracticeLevels     []PracticeLevel
}

// PracticeLevel is for PracticeLevel
type PracticeLevel struct {
	ID                 uuid.UUID `yaml:"id" gorm:"type:uuid;primaryKey"`
	Objective          string    `yaml:"objective"`
	Type               string    `yaml:"type"`
	SecurityPracticeID string    `yaml:"practice" gorm:"type:uuid"`
	MaturityLevelID    string    `yaml:"maturitylevel" gorm:"type:uuid"`
}

// MaturityLevel is for Ma
type MaturityLevel struct {
	ID          uuid.UUID `yaml:"id" gorm:"type:uuid;primaryKey"`
	Number      int       `yaml:"number"`
	Type        string    `yaml:"type"`
	Description string    `yaml:"description" gorm:"type:text"`
}

// ActivityStream
type ActivityStream struct {
	ID                 uuid.UUID   `yaml:"id" gorm:"type:uuid;primaryKey"`
	SecurityPracticeID uuid.UUID   `yaml:"practice"`
	Order              int         `yaml:"order"`
	Type               string      `yaml:"type"`
	Letter             string      `yaml:"letter"`
	Description        string      `yaml:"description" gorm:"type:text"`
	Activities         []*Activity `gorm:"type:uuid;many2many:stream_activities;"`
}

// Activity
type Activity struct {
	ID                uuid.UUID    `yaml:"id" gorm:"type:uuid;primaryKey"`
	ActivityStreamID  uuid.UUID    `yaml:"stream" gorm:"type:uuid"`
	MaturityLevelID   uuid.UUID    `yaml:"level" gorm:"type:uuid"`
	Title             string       `yaml:"title"`
	Benefit           string       `yaml:"benefit"`
	ShortDescription  string       `yaml:"shortDescription"`
	LongDescription   string       `yaml:"longDescription" gorm:"type:text"`
	Results           StringArray  `yaml:"results" gorm:"type:text[]"`
	Metrics           StringArray  `yaml:"metrics" gorm:"type:text[]"`
	Costs             string       `yaml:"costs"`
	Personnel         StringArray  `yaml:"personnel" gorm:"type:text[]"`
	Notes             string       `yaml:"notes"`
	RelatedActivities []*uuid.UUID `yaml:"relatedActivities" gorm:"type:uuid;many2many:related_activities;ForeignKey:id;References:id"`
	Type              string       `yaml:"type"`
	Questions         []Question   `gorm:"type:uuid;many2many:activity_questions;"`
}

// Question
type Question struct {
	ID         uuid.UUID   `yaml:"id" gorm:"type:uuid;primaryKey"`
	ActivityID uuid.UUID   `yaml:"activity" gorm:"type:uuid"`
	Answerset  string      `yaml:"answerset"`
	Text       string      `yaml:"text" gorm:"type:text"`
	Order      int         `yaml:"order"`
	Quality    StringArray `yaml:"quality,omitempty" gorm:"type:text[]"`
	Type       string      `yaml:"type"`
}

type QuestionQuality struct {
}
