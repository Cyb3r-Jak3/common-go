package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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
	fileName = filepath.Clean(fileName)
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

// GetEnvSecret will get either a OS environment variable. If there is no environment variable set it will check to see if a variable with _FILE is set.
// If so then it it will read the secret name as a filepath and return the content
func GetEnvSecret(secretName string) (secret string) {
	secretName = strings.ToUpper(secretName)
	secret = os.Getenv(secretName)
	if secret == "" {
		filePath, isSet := os.LookupEnv(secretName + "_FILE")
		if !isSet {
			return ""
		}
		filePath = filepath.Clean(filePath)
		file, err := ioutil.ReadFile(filePath)
		if os.IsNotExist(err) {
			return ""
		}
		return string(file)
	}
	return
}

// StringSearch checks an array of strings to see if the target string is in it
func StringSearch(target string, array []string) bool {
	i := sort.SearchStrings(array, target)
	if i < len(array) && array[i] == target {
		return true
	}
	return false
}

// FloatSearch checks an array of float64 to see if the target float is in it
func FloatSearch(target float64, array []float64) bool {
	i := sort.SearchFloat64s(array, target)
	if i < len(array) && array[i] == target {
		return true
	}
	return false
}

// IntSearch checks an array of ints to see if the target int is in it
func IntSearch(target int, array []int) bool {
	i := sort.SearchInts(array, target)
	if i < len(array) && array[i] == target {
		return true
	}
	return false
}

//GetEnv checks if the key exists in the environment variables. If yes then returns that value and if not returns default value
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
