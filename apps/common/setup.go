package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type CommonApp struct {
	*core.AppImpl
}

var _ core.App = (*CommonApp)(nil)

// GetName implements core.App.
// Subtle: this method shadows the method (*AppImpl).GetName of CommonApp.AppImpl.
func (c *CommonApp) GetName() string {
	return "common"
}

// Middleware implements core.App.
func (c *CommonApp) Middleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

// Migrate implements core.App.
func (c *CommonApp) Migrate() {
	c.GetDB().AutoMigrate()
}

// Register implements core.App.
func (c *CommonApp) Register(router *gin.RouterGroup) {

}

func NewCommonApp(config *viper.Viper, db *driver.DB) *CommonApp {
	return &CommonApp{
		AppImpl: core.NewAppImpl("common", config, db, []core.Handler{}),
	}
}

var WireSet = wire.NewSet(NewCommonApp)
