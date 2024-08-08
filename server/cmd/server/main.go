package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/database"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/graph"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/pkg/utils"
	"github.com/marveloper-8/real-time-analytics-dashboard-monorepo/service"
)

func main() {
	const defaultPort = "8080"
	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := utils.InitLogger()
	defer logger.Sync()

	influxDB, err := database.NewInfluxDB(
		os.Getenv("INFLUXDB_URL"),
		os.Getenv("INFLUXDB_TOKEN"),
		os.Getenv("INFLUXDB_ORG"),
		os.Getenv("INFLUXDB_BUCKET"),
		logger,
	)
	if err != nil {
		logger.Fatal("Failed to connect to InfluxDB")
	}
	defer influxDB.Close()

	analyticsService := service.NewAnalyticsService(influxDB, logger)

	resolver := &graph.Resolver{
		AnalyticsService: analyticsService,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("connect to http://localhost:" + port + "/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}