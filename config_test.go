package config

import (
	"fmt"
	"os"
	"strconv"
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

func TestWithEnvOverride(t *testing.T) {
	expected := 98765
	_ = os.Setenv("NUM", strconv.Itoa(expected))

	// Nested structure to test.
	type testStruct struct {
		Str string `json:"str"`
		Num int    `json:"num" env:"NUM"`
		Nested struct {
			Str string `json:"str"`
			Num int    `json:"num" env:"NUM"`
		}
	}

	jsonString := "{\"str\":\"foobar\",\"num\":-1111,\"nested\":{\"str\":\"deadbeef\",\"num\":-2222}}"

	ts := new(testStruct)

	err := fromJson([]byte(jsonString), ts)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if ts.Num != 98765 {
		t.Log("expected:", expected, "actual:", ts.Num)
		t.Fail()
	}

	if ts.Nested.Num != 98765 {
		t.Log("expected:", expected, "actual:", ts.Nested.Num)
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
		Port int16 `env:"PORT"`
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
	os.Setenv("PORT", "9000")
	var c = new(AppConfig)
	if err := FromFile("config_example.yaml", c); err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("%+v\n", *c)
	// OUTPUT: {Server:{Port:9000} Database:{Driver:postgres Hostname:localhost Port:5432 Username:postgres Password:dummy Name:my_database}}
}
