package speedprocessor

import "time"

type speedProcessor struct {
}

func NewSpeedSensor() *speedProcessor {
	return &speedProcessor{}
}

func (s *speedProcessor) Reset() {}

func (s *speedProcessor) Update() (updated bool, speed float64, distance float64, duration time.Duration) {
	return false, 0, 0, 0
}
