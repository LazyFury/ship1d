package app

import (
	"yoomall/apps/app/handler"
	"yoomall/core"
	"yoomall/core/driver"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewWireDefaultApp(config *viper.Viper, db *driver.DB,
	dtkHandler *handler.DtkHandler,
	menuHandler *handler.MenuHandler,
	jtkHandler *handler.JtkHandler,
) *DefaultApp {
	return &DefaultApp{
		Config: config,
		App: core.NewApp("default", config, db, []core.Handler{
			menuHandler,
			jtkHandler,
			dtkHandler,
		}),
		AuthMiddlewares: []gin.HandlerFunc{},
	}
}

var handlerSet = wire.NewSet(
	handler.NewDtkHandler,
	handler.NewMenuHandler,
	handler.NewJtkHandler,
)

var WireSet = wire.NewSet(NewWireDefaultApp, handlerSet)
