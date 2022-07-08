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

func TestSpeedPulsesPushThirdTime(t *testing.T) {
	speedo := speedometerDev{
		speedPulses: getSpeedPulsesZeroValue(),
	}
	t0 := time.Now().Add(-time.Second * 2)
	t1 := time.Now().Add(-time.Second)
	t2 := time.Now()
	speedo.pushSpeedPulse(t0)
	speedo.pushSpeedPulse(t1)
	speedo.pushSpeedPulse(t2)
	if speedo.speedPulses[0] != t1 || speedo.speedPulses[1] != t2 {
		t.Errorf("time[0] must be %v and time[1] must be %v", t1, t2)
	}
}

func TestCalcSpeedNoPulse(t *testing.T) {
	speedo := speedometerDev{
		speedPulses: getSpeedPulsesZeroValue(),
	}
	speed := speedo.calcSpeed(time.Now())
	if speed != 0 {
		t.Errorf("speed must be zero if not pulses")
	}
}

func TestCalcSpeedOnePushed(t *testing.T) {
	speedo := speedometerDev{
		speedPulses:  getSpeedPulsesZeroValue(),
		distPerPulse: 0.25,
	}
	speedo.pushSpeedPulse(time.Now())
	speed := speedo.calcSpeed(time.Now())
	if speed != 0 {
		t.Errorf("speed must be zero if only one pulse")
	}
}

func TestCalcSpeedTwoPushedAndCalculetedBelowPrevDur(t *testing.T) {
	speedo := speedometerDev{
		speedPulses:  getSpeedPulsesZeroValue(),
		distPerPulse: 0.25,
	}
	dur := time.Second
	tr := time.Now()
	t0 := tr.Add(-dur)
	t1 := tr
	tcalc := tr.Add(dur / 2)
	speedo.pushSpeedPulse(t0)
	speedo.pushSpeedPulse(t1)
	speed := speedo.calcSpeed(tcalc)
	if speed != speedo.distPerPulse*3.6 {
		t.Errorf("speed must be 0.9, got %f", speed)
	}
}

func TestCalcSpeedTwoPushedAndCalculetedAbovePrevDur(t *testing.T) {
	speedo := speedometerDev{
		speedPulses:  getSpeedPulsesZeroValue(),
		distPerPulse: 0.25,
	}
	dur := time.Second
	tr := time.Now()
	t0 := tr.Add(-dur)
	t1 := tr
	tcalc := t1.Add(t1.Sub(t0) + dur/2)
	speedo.pushSpeedPulse(t0)
	speedo.pushSpeedPulse(t1)
	speed := speedo.calcSpeed(tcalc)
	if speed != speedo.distPerPulse/2.5*3.6 {
		t.Errorf("speed must be 0.36, got %f", speed)
	}
}
