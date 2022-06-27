package pair

import (
	"testing"
)

func TestMutablePair(t *testing.T) {
	mutablePair := NewMutablePair(1, 2)

	if mutablePair.GetLeft() != 1 || mutablePair.GetRight() != 2 {
		t.Errorf("MutablePair work incorrect")
	}

	mutablePair.SetLeft(10)
	mutablePair.SetRight(4)

	if mutablePair.GetLeft() != 10 || mutablePair.GetRight() != 4 {
		t.Errorf("MutablePair work incorrect")
	}
}

func TestImMutablePair(t *testing.T) {
	immutablePair := NewImmutablePair(1, 2)

	if immutablePair.GetLeft() != 1 || immutablePair.GetRight() != 2 {
		t.Errorf("ImmutablePai work incorrect")
	}

	immutablePair.SetLeft(10)
	immutablePair.SetRight(4)

	if immutablePair.GetLeft() != 10 || immutablePair.GetRight() != 4 {
		t.Errorf("ImmutablePai work incorrect")
	}
}
