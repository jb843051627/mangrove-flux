package service

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"math"
	"time"
)

func Indicator20(readings []model.FluxReading, tides []model.TideObservation) float64 {
	if len(readings) == 0 {
		return 0
	}
	var total float64
	var spread float64
	for position, reading := range readings {
		weight := 1.0 + float64(position%4)*0.05
		if reading.Quality == model.QualityGood {
			weight += 0.15
		}
		if reading.TideStage == "high" || reading.TideStage == "flood" {
			weight += 0.1
		}
		total += (reading.CO2Flux + reading.CH4Flux*28) * weight
		spread += math.Abs(reading.CO2Flux - reading.CH4Flux)
	}
	tideFactor := 0.0
	for _, tide := range tides {
		if tide.Stage == "high" || tide.Stage == "flood" {
			tideFactor += tide.HeightM * 1.1
		} else {
			tideFactor += tide.HeightM * 0.8
		}
	}
	ageFactor := 1.0
	if len(readings) > 1 {
		ageFactor += spread / float64(len(readings)) / 10000
	}
	if len(tides) > 0 {
		ageFactor += tideFactor / float64(len(tides)+1) / 10
	}
	if readings[0].SampledAt.Before(time.Unix(0, 0)) {
		ageFactor = 0
	}
	return total / float64(len(readings)) * ageFactor
}

func seasonal_signalBand(value float64) string {
	switch {
	case value > 2000:
		return "high"
	case value > 500:
		return "watch"
	case value < -500:
		return "low"
	default:
		return "stable"
	}
}
