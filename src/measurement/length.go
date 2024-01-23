package measurement

type Length struct {
	Measurement[LengthUnit]
}

func NewLength(value float64, unit LengthUnit) Measurement[LengthUnit] {
	return Measurement[LengthUnit]{value: value, unit: unit}
}
