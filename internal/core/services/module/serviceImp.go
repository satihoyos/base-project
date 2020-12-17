package module

import (
	"sifiv/internal/core/domain"
	port "sifiv/internal/core/ports/module"
)

type moduleImp struct {
}

func NewModuleService() port.ModuleService {
	return &moduleImp{}
}

func (ms *moduleImp) GeAll() []domain.Module {
	modules := []domain.Module{
		domain.Module{
			Code:        "code 1",
			Description: "description 1",
			State:       "state 1",
		},
		domain.Module{
			Code:        "code 2",
			Description: "description 2",
			State:       "state 2",
		},
	}
	return modules
}
