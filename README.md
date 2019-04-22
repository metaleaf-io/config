# config

[![Build Status](https://travis-ci.org/metaleaf-io/config.svg)](https://travis-ci.org/metaleaf-io/config)
[![GoDoc](https://godoc.org/github.com/metaleaf-io/config/github?status.svg)](https://godoc.org/github.com/metaleaf-io/config)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.11.4-green.svg)](https://golang.org/doc/devel/release.html#go1.11)
[![Go version](https://img.shields.io/badge/go-~%3E1.12.0-green.svg)](https://golang.org/doc/devel/release.html#go1.12)

After a search for a *simple* application configuration parser turned up
many complex libraries, I decided to write my own.

## Features

* File formats: JSON and YAML
* Strongly typed configuration data (no key-value maps)
* Data structure is defined by the application.

## Usage

A common use-case for configuration files is for microservice applications.
As an example, let's say we need to configure the port number the web service
listens on and the connection information for a database. The YAML file could
look like this:

```YAML
---
server:
  port: 8080

database:
  driver:   postgres
  hostname: localhost
  port:     5432
  username: postgres
  password: dummy
  name:     my_database
```

Now in the Go application a data structure is defined to match:

```go
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
```

To ingest the configuration file into the application is two lines of code:

```go
var c = new(AppConfig)
var err = config.FromFile("/path/of/config.yaml", c)
```

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Metaleaf.io](http://github.com/metaleaf-io/)

## License

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
