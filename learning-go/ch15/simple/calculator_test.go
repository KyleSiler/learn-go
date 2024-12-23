package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("incorrect result: %d does not match expected %d", result, 3)
	}

	result = Add(1_000, 2_000)

	if result != 3_000 {
		t.Errorf("incorrect result: %d does not match expected %d", result, 3_000)
	}
}

func TestSub(t *testing.T) {
	result := Sub(5, 3)

	if result != 2 {
		t.Errorf("incorrect result: %d does not match expected %d", result, 2)
	}

	result = Sub(4_000, 2_000)

	if result != 2_000 {
		t.Errorf("incorrect result: %d does not match expected %d", result, 2_000)
	}
}
