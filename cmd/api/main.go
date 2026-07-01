package main

import (
	"time"
	"log"
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
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 30 * time.Second,
	}

	//lancer le serveur 
	// Graceful Shutdown
	// crash le programme avec log.Fatal
	// S’il plante pour une vraie raison, affiche l’erreur et quitte.
	// S’il est fermé normalement, ne considère pas ça comme une erreur.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}