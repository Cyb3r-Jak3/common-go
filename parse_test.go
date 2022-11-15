package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

)

type testStruct struct {
	TestString string   `json:"string" yaml:"string"`
	TestInt    int      `json:"int" yaml:"int"`
	TestSlice  []string `json:"slice" yaml:"slice"`
	TestBool   bool     `json:"bool" yaml:"bool"`
}

var expectedStruct = &testStruct{
	TestString: "string",
	TestInt:    5,
	TestSlice:  []string{"string 1", "string 2"},
	TestBool:   false,
}

func TestJSONParse(t *testing.T) {
	testStruct := new(testStruct)
	err := ParseYamlOrJSON("./testData/parsetest.json", testStruct)
	if err != nil {
		t.Errorf("Got an error when reading the test json file. Error: %s", err)
	}
	if !reflect.DeepEqual(expectedStruct, testStruct) {
		t.Errorf("The structs do not match. Expected %v, Actual %v", expectedStruct, testStruct)
	}

}

func TestYAMLParse(t *testing.T) {
	testStruct := new(testStruct)
	err := ParseYamlOrJSON("./testData/parsetest.yml", testStruct)
	if err != nil {
		t.Errorf("Got an error when reading the test yaml file. Error: %s", err)
	}
	if !reflect.DeepEqual(expectedStruct, testStruct) {
		t.Errorf("The structs do not match. Expected %v, Actual %v", expectedStruct, testStruct)
	}

}

func TestBadParse(t *testing.T) {
	testStruct := new(testStruct)
	err := ParseYamlOrJSON("no_file", testStruct)
	if !os.IsNotExist(err) {
		t.Errorf("Error with missing file. Wanted not exists error and got %s", err)
	}
	_ = ioutil.WriteFile("typo.jsno", []byte("test"), 0644)
	err = ParseYamlOrJSON("typo.jsno", testStruct)
	if err.Error() != "unknown file extension for: typo.jsno" {
		t.Errorf("Wanted error with bad file extension. Got %s", err)
	}
	os.Remove("typo.jsno")
}

func TestEnvSecret(t *testing.T) {
	os.Setenv("HELLO", "World")
	secret := GetEnvSecret("Hello")
	if secret != "World" {
		t.Errorf("Error getting Environment Variable. Wanted Hello and got %s", secret)
	}
	secret = GetEnvSecret("World")
	if secret != "" {
		t.Errorf("Error getting invalid environment variable. Wanted blank and got %s", secret)
	}
}

func TestFileSecret(t *testing.T) {
	testData := "SecretSecret"
	err := ioutil.WriteFile("test", []byte(testData), 0644)
	if err != nil {
		t.Errorf("Unable to write testing file. Error: %s", err)
	}
	secret := GetEnvSecret("TEST_FILE")
	if secret != "" {
		t.Errorf("Wanted blank output. Got %s", secret)
	}
	os.Setenv("TEST_FILE", "test")
	secret = GetEnvSecret("test")
	if secret != testData {
		t.Errorf("Error getting secret file. Wanted 'SecretSecret' got %s", secret)
	}
	os.Setenv("TEST_FILE", "emptyfile")
	secret = GetEnvSecret("test")
	if secret != "" {
		t.Errorf("Wanted blank output for missing file. Got %s", secret)
	}
	os.Remove("test")
}

func TestStringSearch(t *testing.T) {
	array := []string{"hello", "world"}
	if !StringSearch("hello", array) {
		t.Errorf("Wanted string in array but it was not found")
	}
	if !StringSearch("world", array) {
		t.Errorf("Wanted string in array but it was not found")
	}
	if StringSearch("fail", array) {
		t.Errorf("Wanted false result and a string was found")
	}
}

func TestFloatSearch(t *testing.T) {
	array := []float64{1.1, 1.2}
	if !FloatSearch(1.1, array) {
		t.Errorf("Wanted flat in array but it was not found")
	}
	if FloatSearch(1.3, array) {
		t.Errorf("Wanted false result and a float was found")
	}
}

func TestIntSearch(t *testing.T) {
	array := []int{1, 2}
	if !IntSearch(1, array) {
		t.Errorf("Wanted int in array but it was not found")
	}
	if IntSearch(3, array) {
		t.Errorf("Wanted false result and a int was found")
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("test", "value")
	returnValue := GetEnv("test", "")
	if returnValue != "value" {
		t.Errorf("Wanted 'value' and got %s", returnValue)
	}
	returnValue = GetEnv("missing", "test")
	if returnValue != "test" {
		t.Errorf("Wanted 'test' and got %s", returnValue)
	}
	os.Unsetenv("test")
}

type KeyValue struct {
	Value string `json:"key"`
}

func TestSkipRoot(t *testing.T) {
	jsonString := `{"root": {"key": "value"}}`
	var Encoded KeyValue
	_ = json.Unmarshal(SkipRoot([]byte(jsonString)), &Encoded)
	if Encoded.Value != "value" {
		t.Errorf("Wanted 'value' and got %s", Encoded.Value)
	}
}

func TestSkipRootMissingRoot(t *testing.T) {
	jsonString := `{"key": "value"}`
	var Encoded KeyValue
	_ = json.Unmarshal(SkipRoot([]byte(jsonString)), &Encoded)

	if Encoded.Value != "" {
		t.Errorf("Wanted '' and got %s", Encoded.Value)
	}
}

func TestSkipRootMissing(t *testing.T) {
	jsonString := ``
	var Encoded KeyValue
	_ = json.Unmarshal(SkipRoot([]byte(jsonString)), &Encoded)
	if Encoded.Value != "" {
		t.Errorf("Wanted '' and got %s", Encoded.Value)
	}
}

func TestSkipRootwithErrorMissing(t *testing.T) {
	jsonString := ``
	_, err := SkipRootwithError([]byte(jsonString))
	if err == nil {
		t.Error("Wanted an error and did not get one")
	}
}

func TestSkipRootwithErrorMissingRoot(t *testing.T) {
	jsonString := `{"key": "value"}`
	value, err := SkipRootwithError([]byte(jsonString))
	if err != nil {
		t.Error("Wanted an error and did not get one")
	}
	var Encoded KeyValue
	_ = json.Unmarshal(value, &Encoded)

	if Encoded.Value != "" {
		t.Errorf("Wanted '' and got %s", Encoded.Value)
	}
}

func TestEnvironMap(t *testing.T) {
	os.Setenv("Test", "value")
	os.Setenv("TestWithEquals", "value=value")
	result := EnvironMap()
	if len(result) == 0 {
		t.Error("Returned map has no length")
	}
	if result["Test"] != "value" {
		t.Errorf("Wanted 'test' and got %s", result["test"])
	}
	if result["TestWithEquals"] != "value=value" {
		t.Errorf("Wanted 'value=value' and got %s", result["TestWithEquals"])
	}
	os.Unsetenv("test")
	os.Setenv("TestWithEquals", "value=value")
}
