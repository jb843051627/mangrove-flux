package model

type Plot struct {
	ID           string  `json:"id"`
	StationID    string  `json:"station_id"`
	Name         string  `json:"name"`
	Habitat      string  `json:"habitat"`
	AreaM2       float64 `json:"area_m2"`
	TargetDepthM float64 `json:"target_depth_m"`
	Active       bool    `json:"active"`
}

func (p Plot) Validate() error {
	if p.ID == "" || p.StationID == "" || p.Name == "" {
		return ErrInvalid
	}
	if p.AreaM2 <= 0 || p.TargetDepthM < 0 {
		return ErrInvalid
	}
	return nil
}
