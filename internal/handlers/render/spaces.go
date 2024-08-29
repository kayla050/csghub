package renderHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SpaceHandler interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
}

type SpaceHandlerImpl struct{}

func NewSpaceHandler() SpaceHandler {
	return &SpaceHandlerImpl{}
}

func (i *SpaceHandlerImpl) List(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "spaces_index", CreateTemplateData(ctx, nil))
}

func (i *SpaceHandlerImpl) Detail(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	spaceName := ctx.Param("space_name")

	ctx.HTML(http.StatusOK, "spaces_show", CreateTemplateData(ctx, map[string]interface{}{
		"namespace": namespace,
		"spaceName": spaceName,
	}))
}