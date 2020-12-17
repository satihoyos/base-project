package module

import "net/http"

type ModuleHandler interface {
	GetAllModules(w http.ResponseWriter, r *http.Request)
}