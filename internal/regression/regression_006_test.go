package regression

import (

	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"

)


func TestBug06_CacheResultsAreIndependent(t *testing.T) {
	cache := service.NewReadingCache(); original := model.FluxReading{DeploymentID: "deployment-06", CO2Flux: 73, SampledAt: time.Now().UTC()}; cache.Update(original)
	values := cache.All(original.DeploymentID); values[0].CO2Flux = 9000
	next := cache.All(original.DeploymentID); if next[0].CO2Flux != original.CO2Flux { t.Fatalf("cache entry changed through returned slice: %.1f", next[0].CO2Flux) }
}

