package main

import (
	"EnployeeService/api"
	"EnployeeService/api/handlers"
	"EnployeeService/config"
	"EnployeeService/dal"
	"EnployeeService/service"
	"context"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	server        *http.Server
	settings      *config.Settings
	mainCtx       context.Context
}

func NewApp(cfg *config.Settings, ctx context.Context) *App {
	return &App{
		server:   nil,
		settings: cfg,
		mainCtx:  ctx,
	}
}

func (a *App)InitServices() {
	settings := config.ReadCfg()
	repos := dal.NewEmployeeRepository(settings.ConnectionString)
	service := service.NewEmployeeService(repos)
	handler := handlers.NewEmployeeHandler(service)
	a.server = api.NewServer(handler, a.mainCtx, a.settings.Port)
}

func (a *App)Start()  {
	go func() {
		log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", a.settings.Port, a.settings.Env))
		a.server.ListenAndServe()
	}()
}


