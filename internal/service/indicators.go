package service

import "github.com/jb843051627/mangrove-flux/internal/model"

func (l *Lab) IndicatorSet(readings []model.FluxReading, tides []model.TideObservation) map[string]float64 {
	return map[string]float64{
		"tidal_exposure": Indicator01(readings, tides), "salinity_stress": Indicator02(readings, tides),
		"methane_quotient": Indicator03(readings, tides), "canopy_respiration": Indicator04(readings, tides),
		"chamber_drift": Indicator05(readings, tides), "closure_stability": Indicator06(readings, tides),
		"waterline_pressure": Indicator07(readings, tides), "carbon_balance": Indicator08(readings, tides),
		"sample_freshness": Indicator09(readings, tides), "plot_coverage": Indicator10(readings, tides),
		"sensor_agreement": Indicator11(readings, tides), "rain_impact": Indicator12(readings, tides),
		"temperature_sensitivity": Indicator13(readings, tides), "flux_variance": Indicator14(readings, tides),
		"goodness_ratio": Indicator15(readings, tides), "deployment_density": Indicator16(readings, tides),
		"night_day_delta": Indicator17(readings, tides), "ebb_flood_skew": Indicator18(readings, tides),
		"quality_backlog": Indicator19(readings, tides), "seasonal_signal": Indicator20(readings, tides),
		"storm_recovery": Indicator21(readings, tides), "station_completeness": Indicator22(readings, tides),
		"chamber_utilization": Indicator23(readings, tides), "carbon_confidence": Indicator24(readings, tides),
	}
}
