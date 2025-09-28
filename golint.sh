#!/bin/bash

cd cmd
go vet
cd ..
cd internal/api
cd ..
go vet
cd config
go vet
cd ..
cd entities
go vet
