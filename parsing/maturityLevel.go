package parsing

import (
	"io/ioutil"
	"log"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessMaturityLevel processes a yaml file to get a markdown
func ProcessMaturityLevel(filename string) (a *model.MaturityLevel, err error) {
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
