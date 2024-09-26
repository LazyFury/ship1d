package handler

import (
	"net/http"

	"lazyfury.github.com/yoomall-server/apps/common/model"
	commonservice "lazyfury.github.com/yoomall-server/apps/common/service"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/helper/curd"
)

type NotFoundRecordHandler struct {
	DB      *driver.DB
	service *commonservice.NotFoundRecordService
	curd    *curd.CRUD
}

// GetRouterGroupName implements core.Handler.
func (n *NotFoundRecordHandler) GetRouterGroupName() string {
	return "notfound"
}

// Register implements core.Handler.
func (n *NotFoundRecordHandler) Register(router *core.RouterGroup) {
	router.WithDoc(&core.DocItem{
		Method: http.MethodGet,
		Path:   "",
	}).GET("", n.curd.GetListHandler(&[]model.NotFoundRecord{}))
}

var _ core.Handler = (*NotFoundRecordHandler)(nil)

func NewNotFoundRecordHandler(db *driver.DB, service *commonservice.NotFoundRecordService) *NotFoundRecordHandler {
	return &NotFoundRecordHandler{
		DB:      db,
		service: service,
		curd: &curd.CRUD{
			DB:    db,
			Model: &model.NotFoundRecord{},
		},
	}
}