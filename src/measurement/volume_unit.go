package measurement

type VolumeUnit string

const (
	LITER      VolumeUnit = "LITER"
	MILLILITER VolumeUnit = "MILLILITER"
)

var literPerUnit = map[VolumeUnit]float64{
	LITER:      1,
	MILLILITER: 0.001,
}

func (w VolumeUnit) getConversionFactor() float64 {
	return literPerUnit[w]
}
