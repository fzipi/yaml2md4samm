package parsing

import (
	"io/ioutil"
	"log"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessPracticeLevel processes a yaml file to get a markdown
func ProcessPracticeLevel(filename string) (pl *model.PracticeLevel, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &pl)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return pl, err
}
