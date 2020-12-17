package main

import (
	"fmt"
	"log"
	"net/http"
	service "sifiv/internal/core/services/module"
	handler "sifiv/internal/infraestructure/web/go/module"
)

func AddModule() {
	serviceModule := service.NewModuleService()
	adapterHandler := handler.NewModuleHandler(serviceModule)
	http.HandleFunc("/modules", adapterHandler.GetAllModules)
}

func AddHome() {
	http.HandleFunc("/", HomePage)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit... prueba cambio")
}

func main() {
	AddModule()
	AddHome()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
