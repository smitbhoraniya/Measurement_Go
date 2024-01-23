package measurement

type LengthUnit string

const (
	METER      LengthUnit = "METER"
	KILOMETERE LengthUnit = "KILOMETER"
	CENTIMETER LengthUnit = "CENTIMETER"
	MILIMETER  LengthUnit = "MILIMETER"
)

var meterPerUnit = map[LengthUnit]float64{
	METER:      1,
	KILOMETERE: 1000,
	CENTIMETER: 0.01,
	MILIMETER:  0.001,
}

func (l LengthUnit) getConversionFactor() float64 {
	return meterPerUnit[l]
}
