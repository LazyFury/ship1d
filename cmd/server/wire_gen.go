// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"lazyfury.github.com/yoomall-server/apps/app"
	"lazyfury.github.com/yoomall-server/apps/app/handler"
	"lazyfury.github.com/yoomall-server/apps/auth"
	handler2 "lazyfury.github.com/yoomall-server/apps/auth/handler"
	"lazyfury.github.com/yoomall-server/apps/auth/middleware"
	"lazyfury.github.com/yoomall-server/apps/auth/service"
	"lazyfury.github.com/yoomall-server/apps/common"
	handler3 "lazyfury.github.com/yoomall-server/apps/common/handler"
	"lazyfury.github.com/yoomall-server/apps/common/service"
	"lazyfury.github.com/yoomall-server/apps/post"
	"lazyfury.github.com/yoomall-server/config"
	"lazyfury.github.com/yoomall-server/core/http"
)

// Injectors from wire.go:

func NewApp() httpserver.HttpServer {
	viper := config.NewConfig()
	db := NewDB(viper)
	dtkHandler := handler.NewDtkHandler(viper)
	authMiddlewareGroup := authmiddleware.NewAuthMiddlewareGroup(db)
	menuHandler := handler.NewMenuHandler(db, authMiddlewareGroup)
	jtkHandler := handler.NewJtkHandler(viper)
	defaultApp := app.NewWireDefaultApp(viper, db, dtkHandler, menuHandler, jtkHandler)
	authService := service.NewAuthService(db)
	userHandler := handler2.NewUserHandler(db, viper, authService, authMiddlewareGroup)
	userRoleHandler := handler2.NewUserRoleHandler(db, authMiddlewareGroup)
	authApp := auth.NewAuthApp(viper, db, authService, userHandler, userRoleHandler)
	postApp := post.NewDefaultApp(viper, db)
	notFoundRecordService := commonservice.NewNotFoundRecordService(db)
	notFoundRecordHandler := handler3.NewNotFoundRecordHandler(db, notFoundRecordService)
	commonApp := common.NewCommonApp(viper, db, notFoundRecordHandler)
	doc := NewDoc()
	httpServer := NewHttpServer(viper, defaultApp, authApp, postApp, commonApp, notFoundRecordService, doc)
	return httpServer
}
