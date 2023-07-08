package touch_test

import (
	"testing"
	"github.com/marksaravi/speedometer-go/touch"
)

func TestTouchConversion(t *testing.T) {
	testCases:=[]struct{
		xt, yt float64
		xs, ys float64
	} {
		{ xt: 512, yt: 512, xs:60, ys: 60 },
		{ xt: 1800, yt: 512, xs:260, ys: 60 },
		{ xt: 512, yt: 1600, xs:60, ys: 180 },
		{ xt: 1800, yt: 1600, xs:260, ys: 180 },

	}

	ax, bx := touch.ConverionFactors(512, 1800, 60, 260)
	ay, by := touch.ConverionFactors(512, 1600, 60, 180)

	for _, tc:=range testCases {
		xs,ys := touch.Convert(tc.xt, tc.yt, ax, bx, ay, by)
		if xs!=tc.xs || ys!=tc.ys {
			t.Errorf("wanted %f, %f, got %f, %f\n", tc.xs, tc.ys, xs, ys)
		}
	}
}