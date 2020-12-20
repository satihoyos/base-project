package module

import (
	"encoding/json"
	"net/http"
	port "sifiv/internal/core/ports/module"
	handler "sifiv/internal/infraestructure/web/adapter/module"
	"sifiv/internal/core/domain"
	"fmt"
)

func (mh *moduleHanlder) GetAllModules(w *http.ResponseWriter, r *http.Request) {
	modules := mh.GeAll()
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(modules)
}

func (mh *moduleHanlder)SaveModule(w *http.ResponseWriter, r *http.Request){
	var module domain.Module
	err := json.NewDecoder(r.Body).Decode(&module)
	if err != nil {
		//Todo buscar manejar errores o una forma mejor de centralizar los errores
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.Save(module)

	if err != nil {
		//Todo buscar manejar errores o una forma mejor de centralizar los errores
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO manejar headers aquí 
	fmt.Fprintf(*w, "Module: %+v", module)
}

type moduleHanlder struct {
	port.ModuleService
}

func NewModuleHandler(service port.ModuleService) handler.ModuleHandler {
	return &moduleHanlder{service}
}
