#!/bin/bash
migrate -path db/migrations -database "mysql://root:rootpass@tcp(127.0.0.1:3306)/combinedprojectdb" up