#!/bin/bash

go get -u github.com/swaggo/swag/cmd/swag
swag init
go get
go install