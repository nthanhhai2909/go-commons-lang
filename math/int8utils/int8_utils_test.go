package int8utils

import "testing"

func TestAbs(t *testing.T) {
	if Abs(-1) != 1 ||
		Abs(0) != 0 ||
		Abs(1) != 1 {
		t.Errorf("Int8#Abs failed")
	}
}

func TestSum(t *testing.T) {
	if Sum(1, 1) != 2 ||
		Sum(0, 0) != 0 ||
		Sum(1, -1) != 0 {
		t.Errorf("Int8#Sum failed")
	}
}

func TestMax(t *testing.T) {
	if Max(1, 1) != 1 ||
		Max(1, 0) != 1 ||
		Max(1, -1) != 1 {
		t.Errorf("Int8#Max failed")
	}
}

func TestMin(t *testing.T) {
	if Min(1, 1) != 1 ||
		Min(1, 0) != 0 ||
		Min(1, -1) != -1 {
		t.Errorf("Int8#Min failed")
	}
}
