package regression

import (

	"testing"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/policy"

)


func TestBug29_RejectedReadingIsExcludedFromReports(t *testing.T) { reading := model.FluxReading{Quality: model.QualityRejected, CO2Flux: 20, CH4Flux: 1}; if policy.IncludeInReport(reading) { t.Fatal("rejected reading was included in report input") } }

