package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
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
// # OR
//
// response := new(exampleStruct)
//
// err := ParseYamlOrJSON("helloworld.yml" response); err != nil { log.Fatal(err)}.
func ParseYamlOrJSON(fileName string, outputInterface interface{}) error {
	fileName = filepath.Clean(fileName)
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	switch filepath.Ext(fileName) {
	case ".json":
		return json.Unmarshal(file, outputInterface)
	case ".yaml", ".yml":
		return yaml.Unmarshal(file, outputInterface)
	default:
		return fmt.Errorf("unknown file extension for: %s", fileName)
	}
}

// GetEnvSecret will get either an OS environment variable. If there is no environment variable set it will check to see if a variable with _FILE is set.
// If so then it will read the secret name as a filepath and return the content.
func GetEnvSecret(secretName string) (secret string) {
	secretName = strings.ToUpper(secretName)
	secret = os.Getenv(secretName)
	if secret == "" {
		filePath, isSet := os.LookupEnv(secretName + "_FILE")
		if !isSet {
			return ""
		}
		filePath = filepath.Clean(filePath)
		file, err := os.ReadFile(filePath)
		if os.IsNotExist(err) {
			return ""
		}
		return string(file)
	}
	return
}

// StringSearch checks an array of strings to see if the target string is in it.
func StringSearch(target string, array []string) bool {
	sort.Strings(array)
	i := sort.SearchStrings(array, target)
	if i < len(array) && array[i] == target {
		return true
	}
	return false
}

// FloatSearch checks an array of float64 to see if the target float is in it.
func FloatSearch(target float64, array []float64) bool {
	sort.Float64s(array)
	i := sort.SearchFloat64s(array, target)
	if i < len(array) && array[i] == target {
		return true
	}
	return false
}

// IntSearch checks an array of integers to see if the target int is in it.
func IntSearch(target int, array []int) bool {
	sort.Ints(array)
	i := sort.SearchInts(array, target)
	if i < len(array) && array[i] == target {
		return true
	}
	return false
}

// GetEnv checks if the key exists in the environment variables. If yes then returns that value and if not returns default value.
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// SkipRootWithError skips the root struct of a JSON message but will return an error. Taken from https://stackoverflow.com/a/20873511.
func SkipRootWithError(jsonBlob []byte) (json.RawMessage, error) {
	var root map[string]json.RawMessage

	if err := json.Unmarshal(jsonBlob, &root); err != nil {
		return nil, err
	}
	for _, v := range root {
		return v, nil
	}
	return nil, nil
}

// SkipRoot skips the root struct of a JSON message but will return nil if an error happens. Taken from https://stackoverflow.com/a/20873511.
func SkipRoot(jsonBlob []byte) (values json.RawMessage) {
	values, _ = SkipRootWithError(jsonBlob)
	return
}

// EnvironMap returns a string map of environment variables.
func EnvironMap() map[string]string {
	results := make(map[string]string)
	for _, x := range os.Environ() {
		item := strings.SplitN(x, "=", 2)
		results[item[0]] = item[1]
	}
	return results
}

// FileExists is a function to check if the file exists at the path.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
	}
	return true
}

// GetDefaultFromEnv checks if the key exists in the environment variables. If yes then returns that value and if not returns default value.
func GetDefaultFromEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ResilientTime Custom time type to handle both RFC3339, no timezone and no fractional seconds.
type ResilientTime struct {
	time.Time
}

func (t *ResilientTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// Remove quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	// Handle null
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	// Try RFC3339 with timezone
	if ts, err := time.Parse(time.RFC3339Nano, s); err == nil {
		t.Time = ts
		return nil
	}
	// Try RFC3339 without timezone
	if ts, err := time.Parse("2006-01-02T15:04:05.999999", s); err == nil {
		t.Time = ts
		return nil
	}
	// Try RFC3339 without fractional seconds
	if ts, err := time.Parse("2006-01-02T15:04:05", s); err == nil {
		t.Time = ts
		return nil
	}
	return fmt.Errorf("cannot parse time: %s", s)
}
