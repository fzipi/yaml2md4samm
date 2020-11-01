package parsing

import (
	"io/ioutil"
	"log"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessQuestions processes a yaml file to get a markdown
func ProcessQuestions(filename string) (a *model.Question, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &a)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return a, err
}
