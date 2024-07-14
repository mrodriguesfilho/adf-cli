package services

import "adf-cli/internal/models"

type InstallStrategy interface {
	Install(installDir, installVersion string, bundle models.Bundle) error
	ServiceName() string
}
