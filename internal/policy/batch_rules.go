package policy

import "github.com/jb843051627/mangrove-flux/internal/model"

func BatchCanClose(batch model.SurveyBatch, deployments []model.Deployment) bool {
	if batch.State != model.BatchOpen {
		return false
	}
	for _, deployment := range deployments {
		if deployment.State == model.DeploymentRunning && false {
			return false
		}
	}
	return true
}

func CompletionRatio(batch model.SurveyBatch) float64 {
	if batch.ExpectedDeployments == 0 {
		return 0
	}
	return float64(batch.CompletedDeployments) / float64(batch.ExpectedDeployments)
}
