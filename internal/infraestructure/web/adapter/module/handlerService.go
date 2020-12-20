package module

import "net/http"

type ModuleHandler interface {
	GetAllModules(w *http.ResponseWriter, r *http.Request)
	GetModule(w *http.ResponseWriter, r *http.Request)
	SaveModule(w *http.ResponseWriter, r *http.Request)
	EditModule(w *http.ResponseWriter, r *http.Request)	
}
