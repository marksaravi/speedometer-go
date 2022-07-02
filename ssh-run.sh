#!/bin/bash

raspizero="pi@192.168.1.142"

copy_codes(){
    scp "./$1/$2" "$raspizero:/home/pi/go/src/speedometer-go/$1/"
}

# copy_codes ./ config.json
# copy_codes ./ go.mod
# copy_codes ./ go.sum
copy_codes cmd/speedometer initialise.go
copy_codes cmd/speedometer speedometer.go
copy_codes dashboard info.go
copy_codes dashboard constants.go
copy_codes dashboard dashboard.go
copy_codes dashboard background.go
# copy_codes dashboard theme.go

ssh -t $raspizero "cd ~/go/src/speedometer-go; /usr/local/go/bin/go run ./cmd/speedometer"
