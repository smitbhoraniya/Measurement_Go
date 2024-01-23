package volume

type Volume struct {
	value float64
	unit  VolumeUnit
}

func NewVolume(value float64, unit VolumeUnit) Volume {
	return Volume{value: value, unit: unit}
}

func (v Volume) convertToTargetUnit(targetUnit VolumeUnit) Volume {
	var convertToTarget float64 = v.value * v.unit.getConversionFactor() / targetUnit.getConversionFactor()

	return Volume{value: convertToTarget, unit: targetUnit}
}

func (l Volume) add(other Volume) Volume {
	var otherVolumeInSameUnit Volume = other.convertToTargetUnit(l.unit)

	return Volume{value: l.value + otherVolumeInSameUnit.value, unit: l.unit}
}

func (l Volume) subtract(other Volume) Volume {
	var otherVolumeInSameUnit Volume = other.convertToTargetUnit(l.unit)

	return Volume{value: l.value - otherVolumeInSameUnit.value, unit: l.unit}
}
