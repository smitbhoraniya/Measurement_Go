package length

type Length struct {
	value float64
	unit  LengthUnit
}

func NewLength(value float64, unit LengthUnit) Length {
	return Length{value: value, unit: unit}
}

func (l Length) convertToTargetUnit(targetUnit LengthUnit) Length {
	var convertToMeter float64 = l.value * l.unit.getConversionFactor()

	var convertToTarget float64 = convertToMeter / targetUnit.getConversionFactor()

	return Length{value: convertToTarget, unit: targetUnit}
}

func (l Length) add(other Length) Length {
	var otherLengthInSameUnit Length = other.convertToTargetUnit(l.unit)

	return Length{value: l.value + otherLengthInSameUnit.value, unit: l.unit}
}

func (l Length) subtract(other Length) Length {
	var otherLengthInSameUnit Length = other.convertToTargetUnit(l.unit)

	return Length{value: l.value - otherLengthInSameUnit.value, unit: l.unit}
}
