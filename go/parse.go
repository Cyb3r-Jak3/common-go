package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// ParseYamlOrJSON will detect if a file is either a JSON or YAML file and marshal it to the provided interface.
// Yaml files can end in .yml or .yaml
//
// Example:
//
// var response exampleStruct
//
// if err := ParseYamlOrJSON("helloworld.json", &response); err != nil { log.Fatal(err)}
//
// OR
//
// response := new(exampleStruct)
//
// err := ParseYamlOrJSON("helloworld.yml" response); err != nil { log.Fatal(err)}
func ParseYamlOrJSON(fileName string, outputInterface interface{}) (err error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	if strings.HasSuffix(fileName, ".json") {
		err = json.Unmarshal(file, outputInterface)
	} else if strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml") {
		err = yaml.Unmarshal(file, outputInterface)
	} else {
		err = fmt.Errorf("unknown file extension for: %s", fileName)
	}
	return err
}

// GetEnvSecret will get either a OS environment variable or if the secretName ends in _FILE it will read the secret name as a filepath.
func GetEnvSecret(secretName string) (secret string) {
	if strings.HasSuffix(secretName, "_FILE") {
		filePath := os.Getenv(secretName)
		file, err := ioutil.ReadFile(filePath)
		if os.IsNotExist(err) {
			return ""
		}
		return string(file)

	}
	return os.Getenv(secretName)
}
