package regression

import (

	"sync"
	"testing"
	"time"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"

)


func TestBug27_ReadingCachePruneIsSafeWithReaders(t *testing.T) { cache := service.NewReadingCache(); reading := model.FluxReading{DeploymentID: "deployment-27", SampledAt: time.Now().UTC()}; for i := 0; i < 10; i++ { cache.Update(reading) }; var wg sync.WaitGroup; for i := 0; i < 6; i++ { wg.Add(1); go func() { defer wg.Done(); for j := 0; j < 60; j++ { cache.PruneBefore(time.Now().Add(-time.Second)); _ = cache.Sorted(reading.DeploymentID) } }() }; wg.Wait() }

