package resize

import (
	"fmt"
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

	fmt.Println(mitchellnetravali(-.1))
	if mitchellnetravali(-.1) != 0.8700555555201976 {
		fmt.Println("fail 1")
		t.Fail()
	}

	fmt.Println(mitchellnetravali(-1.2))
	if mitchellnetravali(-1.2) != -0.014222222215138162 {
		fmt.Println("fail 2")
		t.Fail()
	}

	fmt.Println(mitchellnetravali(12))
	if mitchellnetravali(12) != 0 {
		fmt.Println("fail 2")
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
