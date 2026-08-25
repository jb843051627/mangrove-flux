package model

func CloneReadings(in []FluxReading) []FluxReading {
	out := make([]FluxReading, len(in))
	copy(out, in)
	return out
}

func CloneAlerts(in []QualityAlert) []QualityAlert {
	out := make([]QualityAlert, len(in))
	copy(out, in)
	return out
}

func ClonePoints(in []CalibrationPoint) []CalibrationPoint {
	out := make([]CalibrationPoint, len(in))
	copy(out, in)
	return out
}
