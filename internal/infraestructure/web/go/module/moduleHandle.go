package module

import (
	"encoding/json"
	"net/http"
	port "sifiv/internal/core/ports/module"
	handler "sifiv/internal/infraestructure/web/adapter/module"
)

func (mh *moduleHanlder) GetAllModules(w http.ResponseWriter, r *http.Request) {
	modules := mh.GeAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modules)
}

type moduleHanlder struct {
	port.ModuleService
}

func NewModuleHandler(service port.ModuleService) handler.ModuleHandler {
	return &moduleHanlder{service}
}
