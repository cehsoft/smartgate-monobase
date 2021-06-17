#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o stream-webrtc services/stream/main.go
GOOS=linux GOARCH=amd64 go build -o spitc-manager services/manager/main.go