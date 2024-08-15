package main

import (
	"app-graph/internal/domain/user"
	"app-graph/internal/infrastructure/api/graph"
	"app-graph/internal/infrastructure/database"
	"app-graph/internal/infrastructure/repository"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not found")
	}

	var OriginAllowed = strings.Split(os.Getenv("ORIGIN_ALLOWED"), ",")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := database.ConnectPg(ctx)
	defer database.Close()

	userRepo := repository.NewSQLUserRepository(db)
	userService := user.NewUserService(userRepo)

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   OriginAllowed,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserApp: userService,
	}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
