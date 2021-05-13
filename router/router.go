package router

import (
	"blog/config"
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
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
}
