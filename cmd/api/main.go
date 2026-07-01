package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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


	// Crée un context qui sera annulé quand le programme reçoit un signal d’arrêt.
	// attendre le signal d’arrêt
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()


	go func() {
		log.Println("server started on :8080")

		//lancer le serveur 
		// crash le programme avec log.Fatal
		// S’il plante pour une vraie raison, affiche l’erreur et quitte.
		// S’il est fermé normalement, ne considère pas ça comme une erreur.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Attends jusqu’à ce que le context soit annulé. (recevoir NotifyContext)
	<-ctx.Done()

	log.Println("server shutting down...")

	// Donne maximum 10 secondes au serveur pour s’arrêter proprement. Après ça, on arrête.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// C’est cette ligne qui fait vraiment le graceful shutdown.
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped gracefully")
	
}