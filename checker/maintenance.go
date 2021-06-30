package checker

import (
	"github.com/NordSec/status-checker-go"
	"github.com/NordSec/status-checker-go/maintenance"
)

func NewMaintenanceChecker(client maintenance.Client) status.Checker {
	return &maintenanceChecker{client}
}

type maintenanceChecker struct {
	client maintenance.Client
}

func (mc *maintenanceChecker) Name() string {
	return "maintenance"
}

func (mc *maintenanceChecker) Status() status.Status {
	if enabled, _ := mc.client.IsMaintenanceEnabled(); enabled {
		return status.MAINTENANCE
	}

	return status.OK
}
