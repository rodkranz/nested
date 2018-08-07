[![Build Status](https://travis-ci.org/rodkranz/nested.svg?branch=master)](https://travis-ci.org/rodkranz/nested)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/rodkranz/nested)
 
# NESTED 
===
Simple function to get value from `map[string]interface` without cast every nest field.

## About

Something you need to get value from `map` and this value is an `interface` type, you need to cast this value 
for an type that you know or in nested case you must cast this value and cast for the time that you want.   

## Install

```shell
go get github.com/rodkranz/nested
```

## Import

```go
import (
  "github.com/rodkranz/nested"
)
```

## Test 
To run the project test

```shell
go test -v --cover
```

## Benchmark
To run the benchmark 

```shell
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

Check benchmark in browser                         
```$shell
pprof -http=:8080 cpu.prof
```

To check this result in browser you must have [pprof](https://github.com/google/pprof)

## Example: 

Example os input data 
```go
data := map[string]interface{}{
    "person": map[string]interface{}{
        "name":  "Rodrigo",
        "level": 3,
    },
    "session": map[string]interface{}{
        "token":  "62vsy29v8y4v248v5y97v1e21v35ce97",
        "expire": "2018-08-08T18:00:00Z",
    },
    "images": map[string]interface{}{},
}
```

This is the most common usage how to get the name as string type: 
```go
personValue, ok := data["person"]
if !ok {
    log.Fatal("There is no person")
}

personMap, ok := personValue.(map[string]interface{})
if !ok {
    log.Fatal("cannot cast value to map")
}

nameInterface, ok := personMap["name"]
if !ok {
    log.Fatal("Person has no name")
}

name, ok := personMap["name"].(string)
if !ok {
    log.Fatal("cannot cast value to string")
}

// name is variable with type string 
fmt.Println("Name is ", name)
``` 

If you want to get `string` name of person you can use: 
```go
if name, found := nested.String("person.name", data); !found {
	log.Fatal("cannot find name value as string")
}

fmt.Println("Found name", name)
```

If you want to get `int` level of person you can use:
```go
if level, found := nested.Int("person.level", data); !found {
    log.Fatal("cannot find level value as int")
}
fmt.Println("User is level", level)
```

If you want to get `time.Time` expire of person you can use:
```go
if expire, found := nested.Time("session.expire", data, time2.RFC3339); !found {
	log.Fatal("cannot find expire value as time.Time")
}
fmt.Println("The token will expire at ", expire)

// You can use as 3 parameter a layout of time
if expire, found := nested.Time("session.expire", data, "2006-01-02"); found {
    log.Fatal("cannot find expire value as time.Time with custom layout")
}
fmt.Println("The token will expire at ", expire)
```

If you want to get `interface` images of person you can use:
```go
if images, found := nested.Interface("images", data); found {
	log.Fatal("cannot find images value as interface{}")
}
fmt.Println("Empty map", images)
```
