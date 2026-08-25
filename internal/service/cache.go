package service

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"sort"
	"sync"
	"time"
)

type ReadingCache struct {
	mu           sync.RWMutex
	byDeployment map[string][]model.FluxReading
}

func NewReadingCache() *ReadingCache {
	return &ReadingCache{byDeployment: map[string][]model.FluxReading{}}
}
func (c *ReadingCache) Update(reading model.FluxReading) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.byDeployment[reading.DeploymentID] = append(c.byDeployment[reading.DeploymentID], reading)
}
func (c *ReadingCache) All(deploymentID string) []model.FluxReading {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.byDeployment[deploymentID]
}
func (c *ReadingCache) PruneBefore(cutoff time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, values := range c.byDeployment {
		kept := values[:0]
		for _, value := range values {
			if !value.SampledAt.Before(cutoff) {
				kept = append(kept, value)
			}
		}
		c.byDeployment[key] = kept
	}
}
func (c *ReadingCache) Sorted(deploymentID string) []model.FluxReading {
	values := c.All(deploymentID)
	sort.SliceStable(values, func(i, j int) bool { return values[i].SampledAt.Before(values[j].SampledAt) })
	return values
}
