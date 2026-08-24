package model

type DeploymentState string

const (
	DeploymentPlanned DeploymentState = "planned"
	DeploymentRunning DeploymentState = "running"
	DeploymentClosed  DeploymentState = "closed"
	DeploymentVoid    DeploymentState = "void"
)

type BatchState string

const (
	BatchDraft   BatchState = "draft"
	BatchOpen    BatchState = "open"
	BatchClosed  BatchState = "closed"
	BatchAborted BatchState = "aborted"
)

type QualityState string

const (
	QualityPending  QualityState = "pending"
	QualityGood     QualityState = "good"
	QualityReview   QualityState = "review"
	QualityRejected QualityState = "rejected"
)

type AlertState string

const (
	AlertOpen         AlertState = "open"
	AlertAcknowledged AlertState = "acknowledged"
	AlertCleared      AlertState = "cleared"
)

func CanDeploymentTransition(from, to DeploymentState) bool {
	switch from {
	case DeploymentPlanned:
		return to == DeploymentRunning || to == DeploymentClosed || to == DeploymentVoid
	case DeploymentRunning:
		return to == DeploymentClosed || to == DeploymentVoid
	case DeploymentClosed, DeploymentVoid:
		return false
	default:
		return false
	}
}

func CanBatchTransition(from, to BatchState) bool {
	switch from {
	case BatchDraft:
		return to == BatchOpen || to == BatchAborted
	case BatchOpen:
		return to == BatchClosed || to == BatchAborted
	case BatchClosed, BatchAborted:
		return false
	default:
		return false
	}
}

func CanAlertTransition(from, to AlertState) bool {
	switch from {
	case AlertOpen:
		return to == AlertAcknowledged
	case AlertAcknowledged:
		return to == AlertCleared
	case AlertCleared:
		return false
	default:
		return false
	}
}

func SeverityRank(value string) int {
	switch value {
	case "critical":
		return 3
	case "warning":
		return 2
	case "notice":
		return 1
	default:
		return 0
	}
}
