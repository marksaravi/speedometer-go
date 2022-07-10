package speedometer

import (
	"testing"
	"time"

	"github.com/marksaravi/devices-go/hardware/gpio"
)

type pulseFake struct {
	level gpio.Level
}

func TestReadPulseBothLow(t *testing.T) {
	pulser := pulseFake{
		level: gpio.Low,
	}
	speedo := speedometerDev{
		pulsePinIn:     &pulser,
		prevPulseLevel: gpio.Low,
		pulseCounter:   0,
	}

	speedoInitialConditions := []gpio.Level{gpio.Low, gpio.High}
	pulsePinLevels := []gpio.Level{
		gpio.Low, gpio.Low, gpio.High, gpio.High,
		gpio.Low, gpio.Low, gpio.High, gpio.High,
		gpio.Low, gpio.Low, gpio.High, gpio.High,
		gpio.Low, gpio.Low, gpio.High, gpio.High,
	}
	readWant := [][]bool{
		{
			false, false, false, false,
			true, false, false, false,
			true, false, false, false,
			true, false, false, false,
		},
		{
			true, false, false, false,
			true, false, false, false,
			true, false, false, false,
			true, false, false, false,
		},
	}
	countWant := [][]int64{
		{
			0, 0, 0, 0,
			1, 1, 1, 1,
			2, 2, 2, 2,
			3, 3, 3, 3,
		},
		{
			1, 1, 1, 1,
			2, 2, 2, 2,
			3, 3, 3, 3,
			4, 4, 4, 4,
		},
	}

	for initialLevelIndex := 0; initialLevelIndex < len(speedoInitialConditions); initialLevelIndex++ {
		speedo.prevPulseLevel = speedoInitialConditions[initialLevelIndex]
		speedo.pulseCounter = 0
		for step := 0; step < len(readWant[initialLevelIndex]); step++ {
			pulser.level = pulsePinLevels[step]
			pulsed := speedo.readPulse()
			counter := speedo.pulseCounter
			if pulsed != readWant[initialLevelIndex][step] ||
				counter != countWant[initialLevelIndex][step] {
				t.Errorf("at start level: %v, step: %d, wanted read: %v, counter: %d, but got read: %v, counter: %d",
					speedoInitialConditions[initialLevelIndex],
					step,
					readWant[initialLevelIndex][step],
					countWant[initialLevelIndex][step],
					pulsed,
					counter,
				)
			}
		}
	}
}

func TestSpeedPulsesZeroValue(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
	}
	if !speedo.speedPulseFrom.IsZero() || !speedo.speedPulseTo.IsZero() {
		t.Errorf("both times must be zero at start")
	}
}

func TestSpeedPulsesPushFirstTime(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
	}
	t1 := time.Now()
	speedo.pushSpeedPulse(t1)
	if !speedo.speedPulseFrom.IsZero() || speedo.speedPulseTo.UnixNano() != t1.UnixNano() {
		t.Errorf("time[0] must zero and time[1] must be %v", t1)
	}
}

func TestSpeedPulsesPushSecondTime(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
	}
	t1 := time.Now().Add(-time.Second)
	t2 := time.Now()
	speedo.pushSpeedPulse(t1)
	speedo.pushSpeedPulse(t2)
	if speedo.speedPulseFrom != t1 || speedo.speedPulseTo != t2 {
		t.Errorf("time[0] must be %v and time[1] must be %v", t1, t2)
	}
}

func TestSpeedPulsesPushThirdTime(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
	}
	t0 := time.Now().Add(-time.Second * 2)
	t1 := time.Now().Add(-time.Second)
	t2 := time.Now()
	speedo.pushSpeedPulse(t0)
	speedo.pushSpeedPulse(t1)
	speedo.pushSpeedPulse(t2)
	if speedo.speedPulseFrom != t1 || speedo.speedPulseTo != t2 {
		t.Errorf("time[0] must be %v and time[1] must be %v", t1, t2)
	}
}

func TestCalcSpeedNoPulse(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
	}
	speed := speedo.calcSpeed(time.Now())
	if speed != 0 {
		t.Errorf("speed must be zero if not pulses")
	}
}

func TestCalcSpeedOnePushed(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
		distPerPulse:   0.25,
	}
	speedo.pushSpeedPulse(time.Now())
	speed := speedo.calcSpeed(time.Now())
	if speed != 0 {
		t.Errorf("speed must be zero if only one pulse")
	}
}

func TestCalcSpeedTwoPushedAndCalculetedBelowPrevDur(t *testing.T) {
	speedo := speedometerDev{
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
		distPerPulse:   0.25,
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
		speedPulseFrom: time.Time{},
		speedPulseTo:   time.Time{},
		distPerPulse:   0.25,
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

func (pulser *pulseFake) Read() gpio.Level {
	return pulser.level
}
