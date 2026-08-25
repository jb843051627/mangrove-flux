package regression

import (

	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/model"

)


func TestBug23_AlertCloneDoesNotExposeBackingSlice(t *testing.T) {
	input := []model.QualityAlert{{ID: "alert-23", StationID: "station-23", Code: "quality", Severity: "warning", Message: "x", State: model.AlertOpen, CreatedAt: time.Now().UTC()}}; clone := model.CloneAlerts(input); clone[0].Message = "changed"; if input[0].Message != "x" { t.Fatalf("alert clone changed source slice") }
}

