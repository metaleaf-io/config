package config

import (
	"fmt"
	"testing"
)

type basicList []interface{}

type basicMap struct {
	_integer int     `json:"integer"`
	_float   float64 `json:"float"`
	_string  string  `json:"string"`
	_boolean bool    `json:"boolean"`
}

type testMap struct {
	*basicMap
	_sublist basicList `json:"list"`
	_submap  basicMap  `json:"map"`
}

type testConfig struct {
	_map testMap `json:"map"`
}

func TestFromJsonFile(t *testing.T) {
	var k = new(testConfig)
	var err = FromFile("config_test.json", k)
	if err != nil {
		t.Fail()
	}
}

func TestFromYamlFile(t *testing.T) {
	var k = new(testConfig)
	var err = FromFile("config_test.yaml", k)
	if err != nil {
		t.Fail()
	}
}

func validate(k *testConfig, t *testing.T) {
	if k._map._boolean != true {
		t.Fail()
	}
	if k._map._float != 3.14159 {
		t.Fail()
	}
	if k._map._integer != 1234 {
		t.Fail()
	}
	if k._map._string != "string" {
		t.Fail()
	}
	if k._map._sublist[0] != true {
		t.Fail()
	}
	if k._map._sublist[1] != "string" {
		t.Fail()
	}
	if k._map._sublist[2] != 3.14159 {
		t.Fail()
	}
	if k._map._sublist[3] != 1234 {
		t.Fail()
	}
	if k._map._submap._boolean != true {
		t.Fail()
	}
	if k._map._submap._float != 3.14159 {
		t.Fail()
	}
	if k._map._submap._integer != 1234 {
		t.Fail()
	}
	if k._map._submap._string != "string" {
		t.Fail()
	}
}

// Example structure that matches config_example.yaml.
// NOTE: if the field names match the names used in the YAML file, the json
// struct tags are not necessary.
type AppConfig struct {
	Server struct {
		Port int16
	}
	Database struct {
		Driver   string
		Hostname string
		Port     int16
		Username string
		Password string
		Name     string
	}
}

// Reads the AppConfig data from config_example.yaml and prints it.
func ExampleFromFile() {
	var c = new(AppConfig)
	if err := FromFile("config_example.yaml", c); err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("%+v\n", *c)
	// OUTPUT: {Server:{Port:8080} Database:{Driver:postgres Hostname:localhost Port:5432 Username:postgres Password:dummy Name:my_database}}
}
