package main

import (
	"context"
	httpin "github.com/zineb-b/identity-verification-platform-go/internal/adapter/in/http"
	"github.com/zineb-b/identity-verification-platform-go/internal/bootstrap"
	"github.com/zineb-b/identity-verification-platform-go/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Création du routeur HTTP
	//mux := http.NewServeMux()

	// import my fct handler
	//healthHandler := httpin.NewHealthHandler()

	// Enregister la route dans le router
	// Quand le client appelle GET /health --> exécute la méthode h.Health
	//mux.HandleFunc("GET /health", healthHandler.Health)

	cfg := config.Load()
	container, err := bootstrap.Build(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer container.Close()

	log.Println("application dependencies initialized")

	router := httpin.NewRouter(httpin.Handlers{
		HealthHandler: container.HealthHandler,
		AuthHandler:   container.AuthHandler,
	})

	//creer le serveur http
	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
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
		log.Printf("server started on %s env=%s", cfg.HTTPAddr, cfg.Env)

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
