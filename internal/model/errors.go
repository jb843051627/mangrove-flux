package model

import "errors"

var (
	ErrNotFound    = errors.New("mangrove record not found")
	ErrInvalid     = errors.New("invalid mangrove record")
	ErrConflict    = errors.New("mangrove state conflict")
	ErrCancelled   = errors.New("mangrove operation cancelled")
	ErrCalibration = errors.New("calibration is not usable")
	ErrQuality     = errors.New("flux quality check failed")
	ErrTransaction = errors.New("mangrove transaction failed")
	ErrClosed      = errors.New("mangrove resource is closed")
)

func IsNotFound(err error) bool { return errors.Is(err, ErrNotFound) }
func IsConflict(err error) bool { return errors.Is(err, ErrConflict) }
