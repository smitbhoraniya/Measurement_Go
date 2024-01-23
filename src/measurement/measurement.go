package measurement

type Unit interface {
	getConversionFactor() float64
}

type Measurement[T interface{ Unit }] struct {
	value float64
	unit  T
}

// func NewMeasurement[T interface{ Unit }](value float64, unit T) Measurement[T] {
// 	return Measurement[T]{value: value, unit: unit}
// }

func (l Measurement[T]) convertToTargetUnit(targetUnit T) Measurement[T] {
	var convertToTarget float64 = l.value * l.unit.getConversionFactor() / targetUnit.getConversionFactor()

	return Measurement[T]{value: convertToTarget, unit: targetUnit}
}

func (l Measurement[T]) add(other Measurement[T]) Measurement[T] {
	var otherMeasurementInSameUnit Measurement[T] = other.convertToTargetUnit(l.unit)

	return Measurement[T]{value: l.value + otherMeasurementInSameUnit.value, unit: l.unit}
}

func (l Measurement[T]) subtract(other Measurement[T]) Measurement[T] {
	var otherMeasurementInSameUnit Measurement[T] = other.convertToTargetUnit(l.unit)

	return Measurement[T]{value: l.value - otherMeasurementInSameUnit.value, unit: l.unit}
}
