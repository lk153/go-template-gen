package users

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	IController interface {
		GetUsers(ctx *gin.Context)
		ParseFile(ctx *gin.Context)
	}

	Controller struct {
		svc Service
	}
)

func InitController(r *gin.Engine, svc Service) Controller {
	c := Controller{svc: svc}
	v1 := r.Group("/v1")
	{
		v1.GET("/users", c.GetUsers)
		v1.POST("/parse_file", c.ParseFile)
	}

	return c
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

func (c Controller) ParseFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Get form err: %s", err.Error())
		return
	}

	if err := uploadFile(ctx, file); err != nil {
		ctx.String(http.StatusInternalServerError, "Upload file err: %s", err.Error())
		return
	}

	ctx.String(http.StatusOK, "File %s uploaded successfully", file.Filename)
}

func uploadFile(ctx *gin.Context, file *multipart.FileHeader) (err error) {
	uploadFolder := "./tmp"
	filename := filepath.Base(file.Filename)
	err = ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", uploadFolder, filename))

	return
}
