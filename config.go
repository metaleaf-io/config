/*
   Copyright 2019 Metaleaf.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// The config package provides a simple application configuration system with
// strongly-typed objects. Configuration is read from either a YAML or JSON
// file and returned to the application.
package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// FromFile reads application configuration from a file found at path, then
// returns the decoded data to the application-supplied object, konf.
//
// Errors reading the file, decoding the file, or if the file is of an
// unexpected type are returned.
func FromFile(path string, konf interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	t := strings.ToLower(filepath.Ext(path))
	if t == ".json" {
		return fromJson(buf, konf)
	} else if t == ".yaml" || t == ".yml" {
		return fromYaml(buf, konf)
	} else {
		return fmt.Errorf("unexpected file format: %s", t)
	}
}

// Decodes a JSON file into the configuration object.
func fromJson(buf []byte, konf interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(buf))
	if err := decoder.Decode(&konf); err != nil {
		return fmt.Errorf("config: while decoding JSON: %v", err)
	}

	return fromEnvironment(konf)
}

// Decodes a YAML file into the configuration object.
func fromYaml(buf []byte, konf interface{}) error {
	err := yaml.Unmarshal(buf, &konf)
	if err != nil {
		return fmt.Errorf("config: while decoding YAML: %v", err)
	}

	return fromEnvironment(konf)
}

func fromEnvironment(konf interface{}) error {
	valueOfKonf := reflect.ValueOf(konf).Elem()
	typeOfKonf := reflect.TypeOf(konf).Elem()

	for i := 0; i < valueOfKonf.NumField(); i++  {
		valueOfField := valueOfKonf.Field(i)
		if valueOfField.Kind() == reflect.Struct && valueOfField.CanInterface() {
			// Depth-first search for fields.
			err := fromEnvironment(valueOfField.Addr().Interface())
			if err != nil {
				return err
			}
		}

		// If the field contains an `env` tag, attempt to load the value from the environment.
		typeOfField := typeOfKonf.Field(i)
		tag, ok := typeOfField.Tag.Lookup("env")
		if ok {
			val, ok := os.LookupEnv(tag)
			if ok {
				// Convert the value to match the field and set it.
				switch valueOfField.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					x, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						return err
					}
					valueOfField.SetInt(x)
				case reflect.Bool:
					x, err := strconv.ParseBool(val)
					if err != nil {
						return err
					}
					valueOfField.SetBool(x)
				case reflect.Float32, reflect.Float64:
					x, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return err
					}
					valueOfField.SetFloat(x)
				case reflect.String:
					valueOfField.SetString(val)
				}
			}
		}
	}

	return nil
}