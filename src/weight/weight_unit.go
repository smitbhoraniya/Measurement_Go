package weight

type WeightUnit string

const (
	GRAM     WeightUnit = "GRAM"
	KILOGRAM WeightUnit = "KILOGRAM"
)

var merterPerUnit = map[WeightUnit]float64{
	GRAM:     1,
	KILOGRAM: 1000,
}

func (w WeightUnit) getConversionFactor() float64 {
	return merterPerUnit[w]
}
