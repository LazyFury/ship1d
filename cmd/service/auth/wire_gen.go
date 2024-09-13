// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"lazyfury.github.com/yoomall-server/apps/auth"
	"lazyfury.github.com/yoomall-server/apps/auth/handler"
	"lazyfury.github.com/yoomall-server/apps/auth/service"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/http"
)

// Injectors from wire.go:

func NewApp() httpserver.HttpServer {
	viper := config.NewConfig()
	db := NewDB(viper)
	authService := service.NewAuthService(db)
	userHandler := handler.NewUserHandler(db, viper, authService)
	authApp := auth.NewAuthApp(viper, db, userHandler)
	httpServer := NewHttpServer(viper, authApp)
	return httpServer
}