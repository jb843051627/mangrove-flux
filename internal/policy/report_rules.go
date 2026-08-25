package policy

import "github.com/jb843051627/mangrove-flux/internal/model"

func IncludeInReport(reading model.FluxReading) bool {
	return model.ValidFlux(reading.CO2Flux) && model.ValidFlux(reading.CH4Flux)
}
func GoodRatio(good, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(good) / float64(total)
}
func ReportID(stationID, day string) string { return stationID + ":" + day }
