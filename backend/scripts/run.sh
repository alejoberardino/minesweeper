#!/bin/bash

# Build swagger docs
swag init

# Build API
go build

# Run
./minesweeper