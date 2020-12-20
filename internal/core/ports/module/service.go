package module

import "sifiv/internal/core/domain"

type ModuleService interface {
	GeAll() []domain.Module
	Save(domain.Module) error
	GetByCode(string)domain.Module
	Edit(domain.Module)
}
