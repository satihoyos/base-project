package module

import (
	"sifiv/internal/core/domain"
	port "sifiv/internal/core/ports/module"
)

type moduleImp struct {
	port port.ModuleRepository
}

func NewModuleService(repo port.ModuleRepository) port.ModuleService {
	return &moduleImp{port: repo}
}

func (ms *moduleImp) GeAll() []domain.Module {
	return ms.port.Get()
}

func (ms *moduleImp) GetByCode(code string)domain.Module{
	return ms.port.GetByCode(code)
}

func (ms *moduleImp) Save(module domain.Module) error {
	//TODO validar datos
	//TODO manejar mejor el error se esta ignorando 
	//TODO validar si existe el código en la tabla
	_ = ms.port.Save(module)

	return nil
}
