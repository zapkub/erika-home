#!/usr/bin/env bash
set -e
echo "build playground"
GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/playground ./cmd/playground/main.go

echo "deploy playground"
rsync -avz ./dist/playground pi@192.168.1.55:/home/pi/.erika/bin/playground
ssh pi@192.168.1.55 "/bin/bash -c 'chmod +x /home/pi/.erika/bin/playground'"

echo "running playground...."
ssh pi@192.168.1.55 /home/pi/.erika/bin/playground
