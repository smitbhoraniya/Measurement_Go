package measurement

import (
	"testing"
)

func TestCompareTwoVolumes(t *testing.T) {
	volume1 := NewVolume(10, LITER)
	volume2 := NewVolume(10, LITER)

	if volume1 != volume2 {
		t.Fatalf("Volume are not equals.")
	}
}

func TestConvertVolumeToTargetUnit(t *testing.T) {
	volume1 := NewVolume(1, LITER).convertToTargetUnit(MILLILITER)
	volume2 := NewVolume(1000, MILLILITER)

	if volume1 != volume2 {
		t.Fatalf("Volume are not equals.")
	}
}

func TestConvertVolumeTo(t *testing.T) {
	volume1 := NewVolume(1, LITER).convertToTargetUnit(MILLILITER)
	volume2 := NewVolume(1000, MILLILITER)

	if volume1 != volume2 {
		t.Fatalf("Volume can not ebe converted.")
	}
}

func TestAddVolumes(t *testing.T) {
	volume1 := NewVolume(1, LITER).add(NewVolume(100, MILLILITER))
	volume2 := NewVolume(1.1, LITER)

	if volume1 != volume2 {
		t.Fatalf("Volume can not be added.")
	}
}

func TestSubtractVolumes(t *testing.T) {
	volume1 := NewVolume(1, LITER).subtract(NewVolume(100, MILLILITER))
	volume2 := NewVolume(0.9, LITER)

	if volume1 != volume2 {
		t.Fatalf("Volume can not be subtracted.")
	}
}
