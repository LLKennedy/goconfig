# goconfig

[![GoDoc](https://godoc.org/github.com/LLKennedy/goconfig?status.svg)](https://godoc.org/github.com/LLKennedy/goconfig)
[![Build Status](https://travis-ci.org/disintegration/imaging.svg?branch=master)](https://travis-ci.org/LLKennedy/goconfig)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/LLKennedy/goconfig.svg)
[![Coverage Status](https://coveralls.io/repos/github/LLKennedy/goconfig/badge.svg?branch=master)](https://coveralls.io/github/LLKennedy/goconfig?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/LLKennedy/goconfig)](https://goreportcard.com/report/github.com/LLKennedy/goconfig)
[![Maintainability](https://api.codeclimate.com/v1/badges/fecd47ae6103fd9a3e4d/maintainability)](https://codeclimate.com/github/LLKennedy/goconfig/maintainability)
[![GitHub](https://img.shields.io/github/license/LLKennedy/goconfig.svg)](https://github.com/LLKennedy/goconfig/blob/master/LICENSE)

Tiered default/env/flag/json config loading in Go.

## Installation

`go get "github.com/LLKennedy/goconfig"`

## Forced Standards

This library assumes a few things about how you're loading and storing config. If you're not a fan of these assumptions, feel free to raise an issue, submit a pull request, fork the repository indefinitely, or just use a different config loading library.

### Priority

You don't get to choose the order of priority. Defaults are used to begin, which are overwritten by environment variables, which are overwritten by values from a config file, which are overwritten by runtime executable flags.

### Input Struct

You must provide a pointer to a struct. If you wish to load any fields from JSON or flags, you must tag these fields with JSON tags.

### Environment Variables

Environment variables are in the form APPNAME_FIELDNAME, where APPNAME is the all caps version of the application name passed into the Load function and FIELDNAME is the all caps version of a field in the struct passed into Load.

### Config File Location

Config files are located at %APPDATA%/AppName/config.json on windows, or ~/AppName/config.json on linux and mac. In this case, AppName is the exact case-sensitive name passed into Load.

### Flag Formatting

Flags can only be used to load string and boolean values.

Flags must be in the form of matched string/interface pairs, which can be generated from the ParseArgs function. Flags passed into ParseArgs must be in matched pairs in the form `-key value`. Keys can be preceeded by any number of dashes, all of which will be stripped, except in the case of booleans. Boolean flags must be unpaired and followed immediately by another key so long as that key starts with at least two dashes. In this case the flag will be set to "true" and the next flag will be processed normally.

For example, the following line of args leads to the map below.
`-key1 value1 -key2 --key3 value3 ---key4 value4 -key5 --key6`

```golang
map[string]interface{}{
    "key1": "value1",
    "key2": "true",
    "key3": "value3",
    "key4": "value4",
    "key5": "true",
    "key6": "true",
}
```

## Basic Usage

```golang
import "github.com/LLKennedy/goconfig"

// Use your custom config struct here
type myConfig struct{
    FieldA           string `json:"a"`
    FieldB           string `json:"b"`
    MoreComplexField string `json:"moreComplexString"`
}

...

config := myConfig {
    FieldA: "something",
    FieldB: "something else",
    MoreComplexFiled: "stuff",
}

// Leave nil to use defaults, results written to pointer
err := goconfig.Load(&config, "myApp", nil, nil, nil)
```

## Testing

On windows, the simplest way to test is to use the powershell script.

`./test.ps1`

To emulate the testing which occurs in build pipelines for linux and mac, run the following:

`go test ./... -race`
