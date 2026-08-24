package model

import "time"

type Deployment struct {
	ID        string          `json:"id"`
	BatchID   string          `json:"batch_id"`
	PlotID    string          `json:"plot_id"`
	ChamberID string          `json:"chamber_id"`
	Operator  string          `json:"operator"`
	State     DeploymentState `json:"state"`
	StartedAt time.Time       `json:"started_at"`
	EndedAt   *time.Time      `json:"ended_at,omitempty"`
	Notes     string          `json:"notes"`
}

func (d Deployment) Validate() error {
	if d.ID == "" || d.BatchID == "" || d.PlotID == "" || d.ChamberID == "" {
		return ErrInvalid
	}
	if d.Operator == "" {
		return ErrInvalid
	}
	return nil
}

func (d Deployment) IsActive() bool { return d.State == DeploymentRunning }
