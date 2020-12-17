package module

import "sifiv/internal/core/domain"

type ModuleRepository interface {
	Get() []domain.Module
}
