#!/bin/bash -e

GO_PATH="/usr/local/go/bin/go"

echo "[weather sample] installing..."
$GO_PATH get github.com/eclipse/paho.mqtt.golang
$GO_PATH get github.com/go-ini/ini
$GO_PATH get github.com/jawher/mow.cli
$GO_PATH install snips-platform-lambda-samples/...

echo "[weather sample] running..."
~/go/bin/weather -c ~/go/src/snips-platform-lambda-samples/go/weather/conf.ini