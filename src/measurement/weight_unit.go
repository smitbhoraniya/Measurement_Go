package measurement

type WeightUnit string

const (
	GRAM     WeightUnit = "GRAM"
	KILOGRAM WeightUnit = "KILOGRAM"
)

var gramPerUnit = map[WeightUnit]float64{
	GRAM:     1,
	KILOGRAM: 1000,
}

func (w WeightUnit) getConversionFactor() float64 {
	return gramPerUnit[w]
}
