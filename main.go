package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	service "sifiv/internal/core/services/module"
	"sifiv/internal/infraestructure/sql"
	repo "sifiv/internal/infraestructure/sql/module"
	handler "sifiv/internal/infraestructure/web/go/module"
	handlerModule "sifiv/internal/infraestructure/web/adapter/module"
)

func moduleHandler(w http.ResponseWriter, r *http.Request) {
	connection := sql.NewPostgresConnection()
	repository := repo.NewModuleRepository(connection)
	serviceModule := service.NewModuleService(repository)
	adapterHandler := handler.NewModuleHandler(serviceModule)
	
	switch r.Method {
		case "GET":manageGetURL(&w,r, adapterHandler)
		case "POST":adapterHandler.SaveModule(&w,r)
	}
}

func manageGetURL(w *http.ResponseWriter, r *http.Request,adapterHandler handlerModule.ModuleHandler){	
	fmt.Printf("URL: %s", r.URL.Path)
	fmt.Println("")

	urlDivided := strings.Split(r.URL.Path,"/modules/")
	if len(urlDivided[1]) > 1{
		//TODO get id 
		//TODO MANAGE EDIT		
		adapterHandler.Get(w,r)		
	}else{
		adapterHandler.GetAllModules(w,r)
	}	
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit... prueba cambio")
}


func main() {
	http.HandleFunc("/modules/", moduleHandler)
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
