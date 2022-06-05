#!/bin/bash

 cp ~/go/src/devices-go/hardware/ili9341/ili9341.go /home/pi/go/pkg/mod/github.com/marksaravi/devices-go@v1.1.0/hardware/ili9341/
 cp ~/go/src/devices-go/colors/rgb565/rgb565.go /home/pi/go/pkg/mod/github.com/marksaravi/devices-go@v1.1.0/colors/rgb565/
 cp ~/go/src/devices-go/cmd/test-ili9341/test-ili9341.go /home/pi/go/pkg/mod/github.com/marksaravi/devices-go@v1.1.0/cmd/test-ili9341/

 go run ./cmd/speedometer