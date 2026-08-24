package model

import "time"

type SurveyBatch struct {
	ID                   string     `json:"id"`
	StationID            string     `json:"station_id"`
	Name                 string     `json:"name"`
	State                BatchState `json:"state"`
	StartedAt            time.Time  `json:"started_at"`
	ClosedAt             *time.Time `json:"closed_at,omitempty"`
	ExpectedDeployments  int        `json:"expected_deployments"`
	CompletedDeployments int        `json:"completed_deployments"`
	Revision             int        `json:"revision"`
}

func (b SurveyBatch) Validate() error {
	if b.ID == "" || b.StationID == "" || b.Name == "" || b.StartedAt.IsZero() {
		return ErrInvalid
	}
	if b.ExpectedDeployments < 0 || b.CompletedDeployments < 0 {
		return ErrInvalid
	}
	return nil
}

func (b SurveyBatch) Complete() bool { return b.CompletedDeployments >= b.ExpectedDeployments }
