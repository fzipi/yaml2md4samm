package parsing

import (
	"io/ioutil"
	"log"
	"yaml2md4samm/model"

	"gopkg.in/yaml.v2"
)

// ProcessActivityStream processes a yaml file to get a markdown
func ProcessActivityStream(filename string) (as *model.ActivityStream, err error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &as)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return as, err
}
