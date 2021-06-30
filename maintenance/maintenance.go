package maintenance

type Client interface {
	IsMaintenanceEnabled() (bool, error)
}
