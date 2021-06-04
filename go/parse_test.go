package common_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	common "github.com/Cyb3r-Jak3/common/go"
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
	err := common.ParseYamlOrJSON("../testData/parsetest.json", testStruct)
	if err != nil {
		t.Errorf("Got an error when reading the test json file. Error: %s", err)
	}
	if !reflect.DeepEqual(expectedStruct, testStruct) {
		t.Errorf("The structs do not match. Expected %v, Actual %v", expectedStruct, testStruct)
	}

}

func TestYAMLParse(t *testing.T) {
	testStruct := new(testStruct)
	err := common.ParseYamlOrJSON("../testData/parsetest.yml", testStruct)
	if err != nil {
		t.Errorf("Got an error when reading the test yaml file. Error: %s", err)
	}
	if !reflect.DeepEqual(expectedStruct, testStruct) {
		t.Errorf("The structs do not match. Expected %v, Actual %v", expectedStruct, testStruct)
	}

}

func TestBadParse(t *testing.T) {
	testStruct := new(testStruct)
	err := common.ParseYamlOrJSON("no_file", testStruct)
	if !os.IsNotExist(err) {
		t.Errorf("Error with missing file. Wanted not exists error and got %s", err)
	}
	_ = ioutil.WriteFile("typo.jsno", []byte("test"), 0644)
	err = common.ParseYamlOrJSON("typo.jsno", testStruct)
	if err.Error() != "unknown file extension for: typo.jsno" {
		t.Errorf("Wanted error with bad file extension. Got %s", err)
	}
	os.Remove("typo.jsno")
}

func TestEnvSecret(t *testing.T) {
	os.Setenv("HELLO", "World")
	secret := common.GetEnvSecret("Hello")
	if secret != "World" {
		t.Errorf("Error getting Environment Variable. Wanted Hello and got %s", secret)
	}
	secret = common.GetEnvSecret("World")
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
	secret := common.GetEnvSecret("TEST_FILE")
	if secret != "" {
		t.Errorf("Wanted blank output. Got %s", secret)
	}
	os.Setenv("TEST_FILE", "test")
	secret = common.GetEnvSecret("test")
	if secret != testData {
		t.Errorf("Error getting secret file. Wanted 'SecretSecret' got %s", secret)
	}
	os.Setenv("TEST_FILE", "emptyfile")
	secret = common.GetEnvSecret("test")
	if secret != "" {
		t.Errorf("Wanted blank output for missing file. Got %s", secret)
	}
	os.Remove("test")
}
