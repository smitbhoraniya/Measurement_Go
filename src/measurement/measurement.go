package measurement

type Unit interface {
	getConversionFactor() float64
}

type Measurement[T interface{ Unit }] struct {
	value float64
	unit  T
}

func (m Measurement[T]) convertToTargetUnit(targetUnit T) Measurement[T] {
	var convertToTarget float64 = m.value * m.unit.getConversionFactor() / targetUnit.getConversionFactor()

	return Measurement[T]{value: convertToTarget, unit: targetUnit}
}

func (m Measurement[T]) compare(other Measurement[T]) bool {
	var otherMeasurementInSameUnit Measurement[T] = other.convertToTargetUnit(m.unit)

	return m.value == otherMeasurementInSameUnit.value
}

func (m Measurement[T]) add(other Measurement[T]) Measurement[T] {
	var otherMeasurementInSameUnit Measurement[T] = other.convertToTargetUnit(m.unit)

	return Measurement[T]{value: m.value + otherMeasurementInSameUnit.value, unit: m.unit}
}

func (m Measurement[T]) subtract(other Measurement[T]) Measurement[T] {
	var otherMeasurementInSameUnit Measurement[T] = other.convertToTargetUnit(m.unit)

	return Measurement[T]{value: m.value - otherMeasurementInSameUnit.value, unit: m.unit}
}
