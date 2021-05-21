package router

import (
	"blog/config"
	"blog/log"
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path"
)

var App *gin.Engine

func init() {

	App = gin.New()
	App.Use(middleware.RateLimitBucket())
	// 这个recovery是一个defer
	App.Use(gin.Logger())
	App.Use(gin.Recovery())
	projectPath := config.Settings.ProjectPath

	App.Static("/static", path.Join(projectPath, "frontend/build/static"))
	App.StaticFS("/upload", http.Dir(path.Join(projectPath, "upload")))

	App.StaticFile("/", path.Join(projectPath, "frontend/build/index.html"))
	App.StaticFile("/asset-manifest.json", path.Join(projectPath, "frontend/build/asset-manifest.json"))
	App.StaticFile("/favicon.ico", path.Join(projectPath, "frontend/build/favicon.ico"))
	App.StaticFile("/manifest.json", path.Join(projectPath, "frontend/build/manifest.json"))
	App.StaticFile("/service-worker.js", path.Join(projectPath, "frontend/build/service-worker.js"))
	App.StaticFile("/precache-manifest.15060c2ac3ac7c561cba8f5c4281d90b.js", path.Join(projectPath, "frontend/build/precache-manifest.15060c2ac3ac7c561cba8f5c4281d90b.js"))
	App.StaticFile("/logo.svg", path.Join(projectPath, "frontend/build/logo.svg"))

	App.POST("/upload/:image", UploadImage)
}

func UploadImage(c *gin.Context) {

	log.Logger.Info(c.Request.URL)

	// 32M
	if err := c.Request.ParseMultipartForm(1 << 25); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(200, Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	//if mulForm,err:=c.MultipartForm()
	multiFiles := c.Request.MultipartForm.File

	if v, ok := multiFiles["file"]; ok {
		if len(v) == 0 {
			c.JSON(200, Response{
				State:   1,
				Message: "no file",
			})
			return
		}
		img, err := v[0].Open()
		if err != nil {
			log.Logger.Error(err.Error())
			c.JSON(200, Response{
				State:   1,
				Message: err.Error(),
			})
			return
		}
		name := uuid.NewString()
		saveIo, err := os.Create(config.Settings.ProjectPath + "/upload/" + name + ".jpg")
		if err != nil {
			log.Logger.Error(err.Error())
			c.JSON(200, Response{
				State:   1,
				Message: err.Error(),
			})
			return
		}
		defer saveIo.Close()
		io.Copy(saveIo, img)
		c.JSON(200, Response{
			State:   0,
			Message: "success",
			Data:    "/upload/" + name + ".jpg",
		})
		return
	}

	c.JSON(200, Response{
		State:   1,
		Message: "failed to get file",
	})

}
