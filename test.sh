#!/bin/bash

#go test -covermode=cover -coverprofile=coverage.out
#go tool cover -html=cover.out
#go tool cover -func=coverage.out
#go tool cover -html=coverage.out

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
pprof -http=:8080 cpu.prof

