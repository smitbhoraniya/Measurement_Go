package length

import (
	"testing"
)

func TestCompareTwoLengths(t *testing.T) {
	length1 := Length{value: 10, unit: KILOMETERE}
	length2 := Length{value: 10, unit: KILOMETERE}

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}

func TestConvertTo(t *testing.T) {
	length1 := Length{value: 1, unit: KILOMETERE}.convertToTargetUnit(METER)
	length2 := Length{value: 1000, unit: METER}

	if length1 != length2 {
		t.Fatalf("length are not equals.")
	}
}
