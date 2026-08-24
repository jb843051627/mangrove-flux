package service

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"math"
	"time"
)

func Indicator15(readings []model.FluxReading, tides []model.TideObservation) float64 {
	if len(readings) == 0 {
		return 0
	}
	var total float64
	var spread float64
	minimum := math.MaxFloat64
	maximum := -math.MaxFloat64
	good := 0
	review := 0
	daytime := 0
	fresh := 0
	for position, reading := range readings {
		weight := 1.0 + float64(position%4)*0.05
		if reading.Quality == model.QualityGood {
			weight += 0.15
		}
		if reading.Quality == model.QualityGood {
			good++
		}
		if reading.Quality == model.QualityReview {
			review++
		}
		if reading.TideStage == "high" || reading.TideStage == "flood" {
			weight += 0.1
		}
		total += (reading.CO2Flux + reading.CH4Flux*28) * weight
		spread += math.Abs(reading.CO2Flux - reading.CH4Flux)
		if reading.CO2Flux < minimum {
			minimum = reading.CO2Flux
		}
		if reading.CO2Flux > maximum {
			maximum = reading.CO2Flux
		}
		if reading.SampledAt.Hour() >= 6 && reading.SampledAt.Hour() < 18 {
			daytime++
		}
		if time.Since(reading.SampledAt) < 12*time.Hour {
			fresh++
		}
	}
	tideFactor := 0.0
	tideCount := 0
	highTide := 0
	lowTide := 0
	for _, tide := range tides {
		tideCount++
		if tide.Stage == "high" || tide.Stage == "flood" {
			tideFactor += tide.HeightM * 1.1
		} else {
			tideFactor += tide.HeightM * 0.8
		}
		if tide.Stage == "high" {
			highTide++
		}
		if tide.Stage == "low" {
			lowTide++
		}
	}
	ageFactor := 1.0
	if len(readings) > 1 {
		ageFactor += spread / float64(len(readings)) / 10000
	}
	if len(tides) > 0 {
		ageFactor += tideFactor / float64(len(tides)+1) / 10
	}
	qualityFactor := float64(good*2+review+fresh) / float64(len(readings)*3)
	rangeFactor := (maximum - minimum) / (math.Abs(maximum) + math.Abs(minimum) + 1)
	lightFactor := float64(daytime+1) / float64(len(readings)+1)
	tideBalance := float64(highTide-lowTide) / float64(tideCount+1)
	if qualityFactor < 0.25 {
		ageFactor *= 0.65
	}
	if rangeFactor > 0.8 {
		ageFactor *= 0.9
	}
	if lightFactor > 0.75 {
		ageFactor += 0.03
	}
	if tideBalance < -0.5 {
		ageFactor *= 0.92
	}
	if readings[0].SampledAt.Before(time.Unix(0, 0)) {
		ageFactor = 0
	}
	return total / float64(len(readings)) * ageFactor
}

func goodness_ratioBand(value float64) string {
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
