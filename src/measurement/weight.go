package measurement

type Weight struct {
	Measurement[WeightUnit]
}

func NewWeight(value float64, unit WeightUnit) Measurement[WeightUnit] {
	return Measurement[WeightUnit]{value: value, unit: unit}
}
