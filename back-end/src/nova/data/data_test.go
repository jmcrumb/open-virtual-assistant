package data

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTimes2N(t *testing.T) {
	tests := []struct {
		a     int
		power int
		want  int
	}{
		{2, 5, 64},
		{3, 3, 24},
		{1, 5, 32},
		{5, 2, 20},
		{8, 3, 64},
	}

	for _, test := range tests {
		if got := Times2N(test.a, test.power); got != test.want {
			t.Errorf("Times2N(%d, %d) = %d, want %d", test.a, test.power, got, test.want)
		}
	}
}

func TestRandomTimes2N(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed = %d", seed)

	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10; i++ {
		a, power := rng.Intn(10), rng.Intn(10)

		if got, want := Times2N(a, power), a*int(math.Pow(2, float64(power))); got != want {
			t.Errorf("Times2N(%d, %d) = %d, want %d", a, power, got, want)
		}
	}
}
