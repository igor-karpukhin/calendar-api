package main

import (
	"context"

	"os/signal"

	"os"
	"syscall"

	"time"

	"github.com/gorilla/mux"
	"github.com/igor-karpukhin/calendar-api/candidate"
	"github.com/igor-karpukhin/calendar-api/configuration"
	"github.com/igor-karpukhin/calendar-api/interviewer"
	"github.com/igor-karpukhin/calendar-api/server"
	"github.com/igor-karpukhin/calendar-api/storage"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		return
	}

	//Get application configuration from flags
	c := configuration.NewApplicationConfiguration()
	c.ParseFlags()

	logger.Debug("application started", zap.Any("configuration", c))

	pgConnection, err := storage.NewPostgresConnection(
		c.Postgres.Address,
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.DBName)
	if err != nil {
		logger.Fatal("unable to connect to Postgre DB", zap.Error(err))
	}

	//Make application context
	applicationContext, cancelF := context.WithCancel(context.Background())

	//Subsribe to some system signals to exit HTTP properly
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	//Create application HTTP router
	router := mux.NewRouter().StrictSlash(true)

	//Create candidates repository, controller and build routes
	candidateController := candidate.NewController(candidate.NewPostgresCandidateRepository(pgConnection, logger), logger)
	candidateController.BuildRoutes(router)

	//Create interviewer repository, controller and build routes
	interviewerController := interviewer.NewController(interviewer.NewPostgresInterviewersRepository(pgConnection, logger), logger)
	interviewerController.BuildRoutes(router)

	//Start HTTP server
	appServer := server.NewHTTPServer(c.Host, c.Port, router, logger)
	appServer.Start(applicationContext)
	logger.Info("Application started")
	//Listen to system signals
	for {
		select {
		case <-chSig:
			cancelF()
		case <-time.After(10 * time.Second):
		}
	}
}
