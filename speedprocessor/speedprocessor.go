package speedprocessor

import (
	"time"
	"math/rand"
)

var randomGen *rand.Rand

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
    randomGen = rand.New(s1)
}

type speedProcessor struct {
}

func NewSpeedSensor() *speedProcessor {
	return &speedProcessor{}
}

func (s *speedProcessor) Reset() {}

var lastUpdate time.Time=time.Now()

func (s *speedProcessor) Update() (updated bool, speed float64, distance float64, duration time.Duration) {
	if time.Since(lastUpdate) >= time.Second {
		lastUpdate=time.Now()
		return true, randomGen.Float64()*10, 10,10
	}
	return false, 0, 0, 0
}
