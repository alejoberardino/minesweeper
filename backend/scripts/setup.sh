#!/bin/bash

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/codegangsta/gin
swag init
go get
go install