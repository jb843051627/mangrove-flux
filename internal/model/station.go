package model

import "time"

type FieldStation struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Region    string    `json:"region"`
	Timezone  string    `json:"timezone"`
	TideDatum string    `json:"tide_datum"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

func (s FieldStation) Validate() error {
	if s.ID == "" || s.Name == "" || s.Timezone == "" {
		return ErrInvalid
	}
	if _, err := time.LoadLocation(s.Timezone); err != nil {
		return ErrInvalid
	}
	return nil
}

func (s FieldStation) Location() *time.Location {
	loc, err := time.LoadLocation(s.Timezone)
	if err != nil {
		return time.UTC
	}
	return loc
}
