package parsing

import (
	"io/ioutil"
	"log"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessSecurityPractice processes a yaml file to get a markdown
func ProcessSecurityPractice(filename string) (sp *model.SecurityPractice, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &sp)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return sp, err
}
