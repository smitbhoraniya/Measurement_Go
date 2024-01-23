package measurement

type Volume struct {
	Measurement[VolumeUnit]
}

func NewVolume(value float64, unit VolumeUnit) Measurement[VolumeUnit] {
	return Measurement[VolumeUnit]{value: value, unit: unit}
}
