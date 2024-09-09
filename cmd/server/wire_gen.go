// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"lazyfury.github.com/yoomall-server/apps/app"
	"lazyfury.github.com/yoomall-server/apps/app/handler"
	"lazyfury.github.com/yoomall-server/apps/app/service"
	"lazyfury.github.com/yoomall-server/apps/post"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/http"
)

// Injectors from wire.go:

func NewApp() httpserver.HttpServer {
	viper := config.NewConfig()
	db := NewDB(viper)
	authService := service.NewAuthService(db)
	userHandler := handler.NewUserHandler(db, viper, authService)
	dtkHandler := handler.NewDtkHandler(viper)
	menuHandler := handler.NewMenuHandler()
	jtkHandler := handler.NewJtkHandler(viper)
	defaultApp := app.NewWireDefaultApp(viper, db, userHandler, dtkHandler, menuHandler, jtkHandler)
	postDefaultApp := post.NewDefaultApp(viper, db)
	httpServer := NewHttpServer(viper, defaultApp, postDefaultApp)
	return httpServer
}
