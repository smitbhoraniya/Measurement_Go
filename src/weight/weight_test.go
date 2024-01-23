package weight

import (
	"testing"
)

func TestCompareTwoWeights(t *testing.T) {
	weight1 := Weight{value: 10, unit: KILOGRAM}
	weight2 := Weight{value: 10, unit: KILOGRAM}

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}

func TestConvertToTargetUnit(t *testing.T) {
	weight1 := Weight{value: 1, unit: KILOGRAM}.convertToTargetUnit(GRAM)
	weight2 := Weight{value: 1000, unit: GRAM}

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}

func TestConvertTo(t *testing.T) {
	weight1 := Weight{value: 1, unit: KILOGRAM}.convertToTargetUnit(GRAM)
	weight2 := Weight{value: 1000, unit: GRAM}

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}

func TestAddWeights(t *testing.T) {
	weight1 := Weight{value: 1, unit: KILOGRAM}.add(Weight{value: 100, unit: GRAM})
	weight2 := Weight{value: 1.1, unit: KILOGRAM}

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}

func TestSubtractWeights(t *testing.T) {
	weight1 := Weight{value: 1, unit: KILOGRAM}.subtract(Weight{value: 100, unit: GRAM})
	weight2 := Weight{value: 0.9, unit: KILOGRAM}

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}
