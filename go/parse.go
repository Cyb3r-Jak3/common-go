package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

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
