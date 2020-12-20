package module

import "sifiv/internal/core/domain"

type ModuleRepository interface {
	Get() []domain.Module
	Save(domain.Module) error
	GetByCode(string)domain.Module
	Edit(domain.Module)
}
