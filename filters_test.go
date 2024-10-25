package resize

import (
	"runtime"
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

	if runtime.GOOS == "darwin" {
		if cubic(-1.2) != -0.06400000000000015 {
			t.Fail()
		}
	} else {
		if cubic(-1.2) != -0.06400000000000006 {
			t.Fail()
		}
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

	var negativeOnePointTwo = mitchellnetravali(-1.2)
	if runtime.GOOS == "darwin" {
		if negativeOnePointTwo != -0.014222222215137865 {
			t.Fail()
		}
	} else {
		if mitchellnetravali(-1.2) != -0.014222222215138162 {
			t.Fail()
		}
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

	if sinc(3.885618329942118e-06) != 1 {
		t.Fail()
	}
}

func TestLanczos3(t *testing.T) {
	t.Parallel()

	if lanczos3(0) != 1 {
		t.Fail()
	}

	if lanczos3(6) != 0 {
		t.Fail()
	}
}
