#!/bin/env sh
CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 CC="zig cc -target arm-linux-gnueabihf" CXX="zig c++ -target arm-linux-gnueabihf" go build -trimpath -ldflags="-w -s" -o led main.go
