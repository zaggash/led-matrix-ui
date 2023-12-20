#!/bin/env sh
CGO_ENABLED=1 \
GOOS=linux \
GOARCH=arm \
GOARM=6 \
 CC="zig cc -target arm-linux-gnueabihf -march=arm1176jz_s -mfpu=vfp -mfloat-abi=hard" \
 CXX="zig c++ -target arm-linux-gnueabihf -march=arm1176jzf_s -mfpu=vfp -mfloat-abi=hard" \
  go build -trimpath -ldflags="-w -s" \
    -o matrix-led-ui_armv6 main.go
