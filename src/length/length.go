package length

type Length struct {
	value float64
	unit  LengthUnit
}

func (l Length) convertToTargetUnit(targetUnit LengthUnit) Length {
	var convertToMeter float64
	switch l.unit {
	case METER:
		convertToMeter = l.value
		break
	case KILOMETERE:
		convertToMeter = l.value * 1000
		break
	}

	var convertToTarget float64
	switch targetUnit {
	case METER:
		convertToTarget = convertToMeter
		break
	case KILOMETERE:
		convertToTarget = convertToMeter / 1000
		break
	}

	return Length{value: convertToTarget, unit: targetUnit}
}
