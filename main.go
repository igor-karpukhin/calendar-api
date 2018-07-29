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
	config := configuration.NewApplicationConfiguration()
	config.ParseFlags()

	logger.Debug("application started", zap.Any("configuration", config))

	mongoSession, err := storage.NewMongoConnection(
		config.Mongo.Hosts,
		config.Mongo.DbName,
		config.Mongo.Username,
		config.Mongo.Password)
	if err != nil {
		logger.Error("unable to create mongo storage", zap.Error(err))
		return
	}

	//Make application context
	applicationContext, cancelF := context.WithCancel(context.Background())

	//Subsribe to some system signals to exit HTTP properly
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	//Create application HTTP router
	router := mux.NewRouter().StrictSlash(true)

	//Create candidates repository, controller and build routes
	candidateController := candidate.NewController(candidate.NewMongoRepository(mongoSession, config.Mongo.DbName))
	candidateController.BuildRoutes(router)

	//Create interviewer repository, controller and build routes
	interviewerController := interviewer.NewController(interviewer.NewMongoRepository(mongoSession, config.Mongo.DbName))
	interviewerController.BuildRoutes(router)

	//Start HTTP server
	appServer := server.NewHTTPServer(config.Host, config.Port, router, logger)
	appServer.Start(applicationContext)

	//Listen to system signals
	for {
		select {
		case <-chSig:
			cancelF()
		case <-time.After(10 * time.Second):
		}
	}
}
