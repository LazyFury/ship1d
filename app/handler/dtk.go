package handler

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/core"
	"lazyfury.github.com/yoomall-server/core/constants"
	"lazyfury.github.com/yoomall-server/core/driver"
	"lazyfury.github.com/yoomall-server/core/response"
	"lazyfury.github.com/yoomall-server/libs/dtk"
)

type DtkHandler struct {
	*handler
	dtkClient *dtk.Dtk
}

func NewDtkHandler(app core.App) Handler {
	clent, err := dtk.NewDtkClient(app.GetConfig().DTK)
	if err != nil {
		log.Fatal(err)
	}
	return &DtkHandler{
		handler: &handler{
			App: app,
		},
		dtkClient: clent,
	}
}

func (d *DtkHandler) DB() *driver.DB {
	return constants.DB
}

func (d *DtkHandler) Register(router *gin.RouterGroup) {
	router.GET("", d.dtk)
}

// 大淘客接口 godoc
//
//	@Summary		获取大淘客接口数据
//	@Description	大淘客接口
//	@Tags			/dtk/tb
//	@Accept			json
//	@Produce		json
//	@Param			keyWords	query	string	true	"keyWords"
//	@Param			pageSize	query	string	true	"pageSize"
//	@Param			pageId		query	string	true	"pageId"
//	@Router			/dtk/tb [get]
func (d *DtkHandler) dtk(ctx *gin.Context) {
	var query map[string]string = make(map[string]string)
	ctx.ShouldBindQuery(&query)

	path := query["path"]
	version := query["version"]
	if path == "" || version == "" {
		response.Error(response.ErrBadRequest, "path or version is empty").Done(ctx)
		return
	}

	method := query["method"]
	if method == "" {
		method = http.MethodGet
	}

	delete(query, "method")
	delete(query, "path")
	delete(query, "version")

	log.Info("dtk", "query", query, "url", ctx.Request.URL.Query())
	body, hit := d.dtkClient.RequestWithCache(path, method, version, query)

	var data map[string]any
	err := json.Unmarshal(body, &data)
	if err != nil {
		response.Error(response.ErrInternalError, err.Error()).Done(ctx)
		return
	}

	response.Success(data).WithExtra(map[string]any{
		"hit": hit,
	}).WithExtra(map[string]any{}).Done(ctx)
}
