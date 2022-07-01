#!/bin/bash

raspizero="pi@192.168.1.142"

copy_codes(){
    scp "./$1/$2" "$raspizero:/home/pi/go/src/speedometer-go/$1/"
}

copy_codes cmd/speedometer speedometer.go
copy_codes dashboard dashboard.go
copy_codes dashboard background.go

ssh -t $raspizero "cd ~/go/src/speedometer-go; /usr/local/go/bin/go run ./cmd/speedometer"
