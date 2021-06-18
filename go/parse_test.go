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

func TestStringSearch(t *testing.T) {
	array := []string{"hello", "world"}
	if !common.StringSearch("hello", array) {
		t.Errorf("Wanted string in array but it was not found")
	}
	if common.StringSearch("fail", array) {
		t.Errorf("Wanted false result and a string was found")
	}
}

func TestFloatSearch(t *testing.T) {
	array := []float64{1.1, 1.2}
	if !common.FloatSearch(1.1, array) {
		t.Errorf("Wanted flat in array but it was not found")
	}
	if common.FloatSearch(1.3, array) {
		t.Errorf("Wanted false result and a float was found")
	}
}

func TestIntSearch(t *testing.T) {
	array := []int{1, 2}
	if !common.IntSearch(1, array) {
		t.Errorf("Wanted int in array but it was not found")
	}
	if common.IntSearch(3, array) {
		t.Errorf("Wanted false result and a int was found")
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("test", "value")
	returnValue := common.GetEnv("test", "")
	if returnValue != "value" {
		t.Errorf("Wanted 'value' and got %s", returnValue)
	}
	returnValue = common.GetEnv("missing", "test")
	if returnValue != "test" {
		t.Errorf("Wanted 'test' and got %s", returnValue)
	}
	os.Unsetenv("test")
}
