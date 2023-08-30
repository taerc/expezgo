#!/bin/bash 
go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:123456@tcp(172.10.40.37:3306)/drone_platform" -schema-path "./modules/drone"
