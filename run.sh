#!/bin/sh
fuser -k 9001/tcp
nohup go run app/app.go&
