package measurement

import (
	"testing"
)

func TestCompareTwoWeights(t *testing.T) {
	weight1 := NewWeight(10, KILOGRAM)
	weight2 := NewWeight(10, KILOGRAM)

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}

func TestCompareTwoWeights2(t *testing.T) {
	weight1 := NewWeight(1, KILOGRAM)
	weight2 := NewWeight(1000, GRAM)

	if !weight1.compare(weight2) {
		t.Fatalf("Weight are not equals.")
	}
}

func TestConvertWeightToTargetUnit(t *testing.T) {
	weight1 := NewWeight(1, KILOGRAM).convertToTargetUnit(GRAM)
	weight2 := NewWeight(1000, GRAM)

	if weight1 != weight2 {
		t.Fatalf("Weight are not equals.")
	}
}

func TestConvertWeightTotragetUnit(t *testing.T) {
	weight1 := NewWeight(1, KILOGRAM).convertToTargetUnit(GRAM)
	weight2 := NewWeight(1000, GRAM)

	if weight1 != weight2 {
		t.Fatalf("Weight can not be converted")
	}
}

func TestAddWeights(t *testing.T) {
	weight1 := NewWeight(1, KILOGRAM).add(NewWeight(100, GRAM))
	weight2 := NewWeight(1.1, KILOGRAM)

	if weight1 != weight2 {
		t.Fatalf("Weight can not be added.")
	}
}

func TestSubtractWeights(t *testing.T) {
	weight1 := NewWeight(1, KILOGRAM).subtract(NewWeight(100, GRAM))
	weight2 := NewWeight(0.9, KILOGRAM)

	if weight1 != weight2 {
		t.Fatalf("Weight can not be subtracted.")
	}
}
