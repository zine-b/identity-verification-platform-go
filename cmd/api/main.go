package main

import (
	"net/http"
	httpin "github.com/zineb-b/identity-verification-platform-go/internal/adapter/in/http"
	
)

func main(){
	
	// Création du routeur HTTP
	mux := http.NewServeMux()

	// import my fct handler
	healthHandler := httpin.NewHealthHandler()

	// Enregister la route dans le router 
	// Quand le client appelle GET /health --> exécute la méthode h.Health
	mux.HandleFunc("GET /health", healthHandler.Health)

	//creer le serveur http
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//lancer le serveur 
	server.ListenAndServe()
}