package policy

import "github.com/jb843051627/mangrove-flux/internal/model"

func StationCanReceive(station model.FieldStation) bool { return station.Active && station.ID != "" }
func PlotCanReceive(plot model.Plot) bool               { return plot.Active && plot.AreaM2 > 0 }
func ChamberCanReceive(chamber model.Chamber) bool {
	return chamber.Active && chamber.VolumeL > 0 && chamber.CalibrationID != ""
}
