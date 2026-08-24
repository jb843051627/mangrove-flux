package model

import "time"

type QualityAlert struct {
	ID        string     `json:"id"`
	ReadingID string     `json:"reading_id"`
	StationID string     `json:"station_id"`
	Code      string     `json:"code"`
	Severity  string     `json:"severity"`
	State     AlertState `json:"state"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"created_at"`
	ClearedAt *time.Time `json:"cleared_at,omitempty"`
}

func (a QualityAlert) Validate() error {
	if a.ID == "" || a.StationID == "" || a.Code == "" || a.Message == "" {
		return ErrInvalid
	}
	if SeverityRank(a.Severity) == 0 {
		return ErrInvalid
	}
	return nil
}
