package speedometer

import (
	"testing"
	"time"
)

func TestSpeedPulsesZeroValue(t *testing.T) {
	speedo := speedometerDev{
		speedPulses: getSpeedPulsesZeroValue(),
	}
	if !speedo.speedPulses[0].IsZero() || !speedo.speedPulses[1].IsZero() {
		t.Errorf("both times must be zero at start")
	}
}

func TestSpeedPulsesPushFirstTime(t *testing.T) {
	speedo := speedometerDev{
		speedPulses: getSpeedPulsesZeroValue(),
	}
	t1 := time.Now()
	speedo.pushSpeedPulse(t1)
	if !speedo.speedPulses[0].IsZero() || speedo.speedPulses[1].UnixNano() != t1.UnixNano() {
		t.Errorf("time[0] must zero and time[1] must be %v", t1)
	}
}

func TestSpeedPulsesPushSecondTime(t *testing.T) {
	speedo := speedometerDev{
		speedPulses: getSpeedPulsesZeroValue(),
	}
	t1 := time.Now().Add(-time.Second)
	t2 := time.Now()
	speedo.pushSpeedPulse(t1)
	speedo.pushSpeedPulse(t2)
	if speedo.speedPulses[0] != t1 || speedo.speedPulses[1] != t2 {
		t.Errorf("time[0] must be %v and time[1] must be %v", t1, t2)
	}
}

func TestCalcSpeed(t *testing.T) {

}
