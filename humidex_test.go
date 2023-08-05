package humidex

import (
	"math"
	"testing"
)

func TestHumidex(t *testing.T) {
	humidex := math.Round(Humidex(28, 65))
	if humidex != 36 {
		t.Fatalf("rounded humidex was %v but should have been 36", humidex)
	}
}
