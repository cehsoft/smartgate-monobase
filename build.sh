#!/bin/bash
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o stream-webrtc services/stream/main.go
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o spitc-manager services/manager/main.go