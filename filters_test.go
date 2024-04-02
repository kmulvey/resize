package resize

import (
	"testing"
)

func TestNearest(t *testing.T) {
	t.Parallel()

	if nearest(10) != 0 {
		t.Fail()
	}

	if nearest(.1) != 1 {
		t.Fail()
	}
}

func TestLinear(t *testing.T) {
	t.Parallel()

	if linear(-.1) != .9 {
		t.Fail()
	}

	if linear(10) != 0 {
		t.Fail()
	}
}

func TestCubic(t *testing.T) {
	t.Parallel()

	if cubic(-.1) != 0.9765 {
		t.Fail()
	}

	if cubic(-1.2) != -0.06400000000000006 {
		t.Fail()
	}

	if cubic(12) != 0 {
		t.Fail()
	}
}

func TestMitchellnetravali(t *testing.T) {
	t.Parallel()

	if mitchellnetravali(-.1) != 0.8700555555201976 {
		t.Fail()
	}

	if mitchellnetravali(-1.2) != -0.014222222215138162 {
		t.Fail()
	}

	if mitchellnetravali(12) != 0 {
		t.Fail()
	}
}

func TestSinc(t *testing.T) {
	t.Parallel()

	if sinc(0.477464829275686) != 0.6649966577360363 {
		t.Fail()
	}

	if sinc(0.35014087480216977) != 0.8101885091467594 {
		t.Fail()
	}
}
