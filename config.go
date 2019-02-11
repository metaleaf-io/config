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
	"path/filepath"
	"strings"
)

// FromFile reads application configuration from a file found at path, then
// returns the decoded data to the application-supplied object, konf.
//
// Errors reading the file, decoding the file, or if the file is of an
// unexpected type are returned.
func FromFile(path string, konf interface{}) error {
	t := strings.ToLower(filepath.Ext(path))
	if t == ".json" {
		return fromJson(path, konf)
	} else if t == ".yaml" || t == ".yml" {
		return fromYaml(path, konf)
	} else {
		return fmt.Errorf("unexpected file format: %s", t)
	}
}

// Decodes a JSON file into the configuration object.
func fromJson(path string, konf interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewReader(buf))
	if err := decoder.Decode(&konf); err != nil {
		return fmt.Errorf("while decoding JSON: %v", err)
	}

	return nil
}

// Decodes a YAML file into the configuration object.
func fromYaml(path string, konf interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &konf)
	if err != nil {
		return err
	}

	return nil
}
