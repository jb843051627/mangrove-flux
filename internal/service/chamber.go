package service

import (
	"context"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"
)

func (l *Lab) RegisterChamber(ctx context.Context, chamber model.Chamber) error {
	if err := chamber.Validate(); err != nil {
		return err
	}
	if _, err := l.requireStation(ctx, chamber.StationID); err != nil {
		return err
	}
	chamber.Active = true
	return l.chambers.Save(ctx, &chamber)
}
func (l *Lab) GetChamber(ctx context.Context, id string) (*model.Chamber, error) {
	return l.chambers.Get(ctx, id)
}
func (l *Lab) ListChambers(ctx context.Context, stationID string) ([]model.Chamber, error) {
	if _, err := l.requireStation(ctx, stationID); err != nil {
		return nil, err
	}
	return l.chambers.ListByStation(ctx, stationID)
}
func (l *Lab) ActivateChamber(ctx context.Context, id, calibrationID string) error {
	chamber, err := l.chambers.Get(ctx, id)
	if err != nil {
		return err
	}
	if calibrationID == "" {
		return model.ErrCalibration
	}
	chamber.CalibrationID = calibrationID
	chamber.Active = true
	if !policy.ChamberCanReceive(*chamber) {
		return model.ErrCalibration
	}
	return l.chambers.Save(ctx, chamber)
}
