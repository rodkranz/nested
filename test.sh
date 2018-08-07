#!/bin/bash

go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out

#go test -covermode=cover -coverprofile=cover.out
#go tool cover -html=cover.out