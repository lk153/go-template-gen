package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IController interface {
		GetUsers(ctx *gin.Context)
	}

	Controller struct {
		svc Service
	}
)

func InitController(svc Service) Controller {
	return Controller{svc: svc}
}

func (c Controller) GetUsers(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	user, err := c.svc.GetUsers(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "impossible to retrieve user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
