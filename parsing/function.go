package parsing

import (
	"io/ioutil"
	"log"

	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessBusinessFunction processes a yaml file to get a markdown
func ProcessBusinessFunction(filename string) (f *model.BusinessFunction, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &f)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return f, err
}
