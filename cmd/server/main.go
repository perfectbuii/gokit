package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	deliveries "github.com/perfectbuii/gokit/internal/deliveries/grpc"
	"github.com/perfectbuii/gokit/internal/repositories"
	"github.com/perfectbuii/gokit/internal/services"
	"github.com/perfectbuii/gokit/pb"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type server struct {

	// repo
	hubRepo repositories.HubRepository

	// service
	hubSrv services.HubService

	// deliveries
	hubpb pb.HubServiceServer
	// other
	db *sql.DB
}

var srv server

func start() error {
	ctx := context.Background()
	if err := srv.loadConfig(); err != nil {
		return err
	}

	if err := srv.loadDB(); err != nil {
		return err
	}

	if err := srv.loadRepositories(); err != nil {
		return err
	}

	if err := srv.loadServices(); err != nil {
		return err
	}

	if err := srv.loadDeliveries(); err != nil {
		return err
	}

	srv.startServer(ctx)
	return nil
}

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) loadDB() error {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *server) loadRepositories() error {
	s.hubRepo = repositories.NewHubRepository(s.db)

	return nil
}

func (s *server) loadServices() error {
	s.hubSrv = services.NewHubService(s.hubRepo)

	return nil
}

func (s *server) loadDeliveries() error {
	s.hubpb = deliveries.NewHubDelivery(s.hubSrv)

	return nil
}

func (s *server) loadConfig() error {
	return nil
}

func (s *server) startServer(ctx context.Context) error {
	var serverError = make(chan error)
	var waitGroup sync.WaitGroup

	httpPort := "8585"
	grpcPort := "5000"

	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		mux := runtime.NewServeMux(
		// runtime.WithMetadata(metadata.Authentication),
		)

		if err := pb.RegisterHubServiceHandlerServer(ctx, mux, s.hubpb); err != nil {
			serverError <- err
		}

		handler := cors.AllowAll().Handler(mux)
		log.Printf("HTTP Server listens on port: %s\n", httpPort)
		http.ListenAndServe(fmt.Sprintf(":%s", httpPort), handler)
	}()

	go func() {
		defer waitGroup.Done()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			serverError <- err
		}

		grpcServer := grpc.NewServer()
		pb.RegisterHubServiceServer(grpcServer, s.hubpb)

		log.Printf("GRPC Server listens on port: %v", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			serverError <- err
		}
	}()

	for <-serverError != nil {
		log.Fatal("start server failed:", <-serverError)
	}

	waitGroup.Wait()

	return nil
}
