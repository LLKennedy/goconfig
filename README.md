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

defaults := myConfig {
    FieldA: "something",
    FieldB: "something else",
    MoreComplexFiled: "stuff",
}

// Leave nil to use defaults
loaded, err := goconfig.Load(defaults, nil, nil, nil)
if err != nil {
    // handle errors
}
filledConfig, _ := loaded.(myConfig) // returned value will ALWAYS be of provided type
```

## Testing

On windows, the simplest way to test is to use the powershell script.

`./test.ps1`

To emulate the testing which occurs in build pipelines for linux and mac, run the following:

`go test ./... -race`
