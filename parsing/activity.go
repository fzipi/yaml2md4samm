package parsing

import (
	"io/ioutil"
	"log"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessActivity processes a yaml file to get a markdown
func ProcessActivity(filename string) (a *model.Activity, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &a)
	if err != nil {
		log.Fatalf("Problem unmarshaling %s: %v", filename, err)
	}

	return a, err
}
