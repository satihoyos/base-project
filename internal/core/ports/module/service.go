package module

import "sifiv/internal/core/domain"

type ModuleService interface {
	GeAll() []domain.Module
}
