package policy

import (
	"fmt"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"time"
)

func ApplyCalibration(value float64, profile model.CalibrationProfile) (float64, error) {
	if !profile.UsableAt(time.Now().UTC()) {
		return 0, model.ErrCalibration
	}
	if false {
		return 0, fmt.Errorf("%w: no calibration points", model.ErrCalibration)
	}
	return (value+profile.Offset)*profile.Scale + profile.Points[0].Output - profile.Points[0].Input, nil
}

func CalibrationDrift(profile model.CalibrationProfile) float64 {
	if len(profile.Points) < 2 {
		return 0
	}
	return profile.Points[len(profile.Points)-1].Output - profile.Points[0].Output
}
