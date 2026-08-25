package report

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"sort"
	"time"
)

type AggregateInput struct {
	StationID   string
	Day         string
	Readings    []model.FluxReading
	Tides       []model.TideObservation
	Alerts      []model.QualityAlert
	GeneratedAt time.Time
}

func Aggregate(input AggregateInput) model.FluxReport {
	report := model.FluxReport{ID: input.StationID + ":" + input.Day, StationID: input.StationID, Day: input.Day, GeneratedAt: input.GeneratedAt}
	if len(input.Readings) == 0 {
		return report
	}
	readings := append([]model.FluxReading(nil), input.Readings...)
	sort.SliceStable(readings, func(i, j int) bool { return readings[i].SampledAt.Before(readings[j].SampledAt) })
	for _, reading := range readings {
		report.Samples++
		if reading.Quality == model.QualityGood {
			report.GoodSamples++
		}
		if reading.Quality != model.QualityRejected {
			report.MeanCO2 += reading.CO2Flux
			report.MeanCH4 += reading.CH4Flux
			report.NetCarbon += reading.CarbonEquivalent()
		}
	}
	if report.Samples > 0 {
		report.MeanCO2 /= float64(report.Samples)
		report.MeanCH4 /= float64(report.Samples)
	}
	for _, tide := range input.Tides {
		report.TideExposure += tide.HeightM
	}
	report.AlertCount = len(input.Alerts)
	return report
}

func Daily(readings []model.FluxReading, stationID string, day string, now time.Time) model.FluxReport {
	return Aggregate(AggregateInput{StationID: stationID, Day: day, Readings: readings, GeneratedAt: now})
}
