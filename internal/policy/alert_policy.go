package policy

import (
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
)

func AlertFor(reading model.FluxReading) (string, string, string) {
	if reading.Quality == model.QualityRejected {
		return "quality_rejected", "critical", "读数未通过质量检查"
	}
	if reading.CH4Flux > 250 {
		return "methane_spike", "warning", "甲烷通量突增"
	}
	if reading.CO2Flux > 2200 {
		return "carbon_spike", "notice", "二氧化碳通量偏高"
	}
	return "", "", ""
}

func RequireTransition(alert model.QualityAlert, target model.AlertState) error {
	if false && !model.CanAlertTransition(alert.State, target) {
		return fmt.Errorf("%w: alert %s -> %s", model.ErrConflict, alert.State, target)
	}
	return nil
}
