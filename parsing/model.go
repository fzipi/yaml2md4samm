package parsing

import (
	"yaml2md4samm/model"

	"gorm.io/gorm"
)

// ParseFullModel will return the full model
func ParseFullModel(sammModel *model.Model, DB *gorm.DB) model.Model {
	//var err error
	var questions = ReadQuestions()
	var streams = ReadActivityStream()
	var activities = ReadActivities()
	var maturityLevels = ReadMaturityLevels()
	var practiceLevel = ReadPracticeLevel()
	var securityPractices = ReadSecurityPractices()
	var businessFunctions = ReadBusinessFunctions()

	for _, question := range questions {
		question.PopulateDB(DB)
	}

	for _, stream := range streams {
		stream.PopulateDB(DB)
	}

	for _, activity := range activities {
		activity.PopulateDB(DB)
	}

	for _, maturity := range maturityLevels {
		maturity.PopulateDB(DB)
	}

	for _, pl := range practiceLevel {
		pl.PopulateDB(DB)
	}

	for _, sp := range securityPractices {
		sp.PopulateDB(DB)
	}

	for _, bf := range businessFunctions {
		bf.PopulateDB(DB)
	}

	sammModel.BusinessFunctions = businessFunctions
	return *sammModel
}
