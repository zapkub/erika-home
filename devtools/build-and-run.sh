#!/usr/bin/env bash
set -e
echo "build erika"
GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/erika ./cmd/erika/main.go

echo "deploy erika"
rsync -avz ./dist/erika pi@192.168.1.55:/home/pi/.erika/bin/erika
ssh pi@192.168.1.55 "/bin/bash -c 'chmod +x /home/pi/.erika/bin/erika'"

echo "running erika...."
ssh pi@192.168.1.55 '/bin/bash -c "kill -9 `pidof erika` && /home/pi/.erika/bin/erika"'
