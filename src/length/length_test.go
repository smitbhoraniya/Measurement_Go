package length

import (
	"testing"
)

func TestCompareTwoLengths(t *testing.T) {
	length1 := NewLength(10, KILOMETERE)
	length2 := NewLength(10, KILOMETERE)

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}

func TestConvertToTargetUnit(t *testing.T) {
	length1 := NewLength(1, KILOMETERE).convertToTargetUnit(METER)
	length2 := NewLength(1000, METER)

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}

func TestConvertTo(t *testing.T) {
	length1 := NewLength(1, KILOMETERE).convertToTargetUnit(METER)
	length2 := NewLength(1000, METER)

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}

func TestAddLengths(t *testing.T) {
	length1 := NewLength(1, KILOMETERE).add(Length{value: 100, unit: METER})
	length2 := NewLength(1.1, KILOMETERE)

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}

func TestSubtractLengths(t *testing.T) {
	length1 := NewLength(1, KILOMETERE).subtract(Length{value: 100, unit: METER})
	length2 := NewLength(0.9, KILOMETERE)

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}
