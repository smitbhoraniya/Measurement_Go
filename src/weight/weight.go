package weight

type Weight struct {
	value float64
	unit  WeightUnit
}

func NewWeight(value float64, unit WeightUnit) Weight {
	return Weight{value: value, unit: unit}
}

func (w Weight) convertToTargetUnit(targetUnit WeightUnit) Weight {
	var convertToTarget float64 = w.value * w.unit.getConversionFactor() / targetUnit.getConversionFactor()

	return Weight{value: convertToTarget, unit: targetUnit}
}

func (l Weight) add(other Weight) Weight {
	var otherWeightInSameUnit Weight = other.convertToTargetUnit(l.unit)

	return Weight{value: l.value + otherWeightInSameUnit.value, unit: l.unit}
}

func (l Weight) subtract(other Weight) Weight {
	var otherWeightInSameUnit Weight = other.convertToTargetUnit(l.unit)

	return Weight{value: l.value - otherWeightInSameUnit.value, unit: l.unit}
}
