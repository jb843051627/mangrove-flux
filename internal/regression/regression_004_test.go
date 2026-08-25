package regression

import (

	"sync"
	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"

)


func TestBug04_ReadingCacheConcurrentUpdateAndRead(t *testing.T) {
	cache := service.NewReadingCache(); reading := model.FluxReading{DeploymentID: "deployment-04", SampledAt: time.Now().UTC()}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ { wg.Add(1); go func() { defer wg.Done(); for j := 0; j < 80; j++ { cache.Update(reading); _ = cache.All(reading.DeploymentID) } }() }
	wg.Wait()
}

