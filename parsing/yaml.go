package parsing

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// Get File names based on a directory and glob pattern
func getFiles(directory string, globPattern string) []string {
	matches, _ := filepath.Glob(directory + "/" + globPattern)

	return matches
}

// ReadBusinessFunctions read the yaml
func ReadBusinessFunctions() []model.BusinessFunction {
	// Globbing patterns for functions:
	functionGlob := "Function*.yml"
	var functions = getFiles("Datafiles", functionGlob)

	var businessFunctions []model.BusinessFunction
	for _, function := range functions {
		bf, err := ProcessBusinessFunction(function)
		if err != nil {
			log.Fatalf("Fatal error parsing %s: %v\n", function, err)
		}
		//log.Printf("Found Business Funtion %s\n", bf.Name)
		businessFunctions = append(businessFunctions, *bf)
	}
	return businessFunctions
}

// ReadSecurityPractices reads the yaml
func ReadSecurityPractices() []model.SecurityPractice {
	// Globbing pattern for practices:
	practiceGlob := "Practice\\ ?-*.yml"
	var practiceFiles = getFiles("Datafiles", practiceGlob)

	var practices []model.SecurityPractice
	for _, practice := range practiceFiles {
		p, err := ProcessSecurityPractice(practice)
		if err != nil {
			log.Fatalf("Fatal error parsing %s: %v\n", practice, err)
		}
		//log.Printf("Found Security Practice %s with id %s\n", p.Name, p.ID)
		practices = append(practices, *p)
	}

	return practices
}

// ReadPracticeLevel reads the yaml
func ReadPracticeLevel() []model.PracticeLevel {
	practiceLevelGlob := "Practice\\ Level*.yml"
	var practicesLevelFiles = getFiles("Datafiles", practiceLevelGlob)
	var practiceLevels []model.PracticeLevel
	for _, plevel := range practicesLevelFiles {
		pls, err := ProcessPracticeLevel(plevel)
		if err != nil {
			log.Fatalf("Fatal error parsing %s: %v\n", plevel, err)
		}
		// log.Printf("Found Practice Level id %s, pointing to %s, with maturity level %+v\n",
		// 	pls.ID,
		// 	pls.SecurityPracticeID,
		// 	pls.MaturityLevelID)
		practiceLevels = append(practiceLevels, *pls)
	}
	return practiceLevels
}

//ReadMaturityLevels reads the yaml
func ReadMaturityLevels() []model.MaturityLevel {
	// Globbing patterns for functions:
	mlGlob := "Maturity\\ Level\\ ?.yml"
	var maturityLevelFiles = getFiles("Datafiles", mlGlob)

	var maturityLevels []model.MaturityLevel
	for _, maturityLevel := range maturityLevelFiles {
		ml, err := ProcessMaturityLevel(maturityLevel)
		if err != nil {
			log.Fatalf("Fatal error parsing %s: %v\n", maturityLevel, err)
		}
		//log.Printf("Found Maturity Level %d, with description %s\n", ml.Number, ml.Description)
		maturityLevels = append(maturityLevels, *ml)
	}
	return maturityLevels
}

// ReadQuestions reads the yaml
func ReadQuestions() []model.Question {
	// Globbing patterns for questions:
	qGlob := "Question\\ *.yml"
	var questionFiles = getFiles("Datafiles", qGlob)

	var questions []model.Question
	for _, question := range questionFiles {
		q, err := ProcessQuestions(question)
		if err != nil {
			log.Fatalf("Fatal error parsing %s: %v\n", question, err)
		}
		//log.Printf("Found Question %s\n", q.ID)
		questions = append(questions, *q)
	}
	return questions
}

// ReadActivities reads the yaml
func ReadActivities() []model.Activity {
	// Globbing patterns for questions:
	aGlob := "Activity\\ *.yml"
	var activityFiles = getFiles("Datafiles", aGlob)

	var activities []model.Activity
	for _, activity := range activityFiles {
		//log.Printf("Parsing activity file %s\n", activity)
		a, err := ProcessActivity(activity)
		if err != nil {
			log.Fatalf("Error parsing %s file, %v\n", activity, err)
		}
		//log.Printf("Found Activity %s\n", a.ID)
		activities = append(activities, *a)
	}
	return activities
}

// ReadActivityStream reads the yaml
func ReadActivityStream() []model.ActivityStream {
	// Globbing patterns for questions:
	asGlob := "Stream\\ *.yml"
	var activityStreamFiles = getFiles("Datafiles", asGlob)

	var activityStreams []model.ActivityStream
	for _, activityStream := range activityStreamFiles {
		as, err := ProcessActivityStream(activityStream)
		if err != nil {
			log.Fatalf("Fatal error parsing %s: %v\n", activityStream, err)
		}
		//log.Printf("Found ActivityStream %s\n", as.ID)
		activityStreams = append(activityStreams, *as)
	}
	return activityStreams
}

// ReadModel processes a yaml file to get a markdown
func ReadModel(filename string) (m *model.Model, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return m, err
}
