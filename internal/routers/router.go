package routers

import (
	"net/http"
	"time"

	_ "github.com/chen2eric/blog-service/docs"
	"github.com/chen2eric/blog-service/global"
	"github.com/chen2eric/blog-service/internal/middleware"
	"github.com/chen2eric/blog-service/internal/routers/api"
	v1 "github.com/chen2eric/blog-service/internal/routers/api/v1"
	"github.com/chen2eric/blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
		r.Use(middleware.Translations())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeOut(global.AppSetting.DefaultConTextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())

	apiv1.POST("/tags", tag.Create)
	apiv1.DELETE("/tags/:id", tag.Delete)
	apiv1.PUT("/tags/:id", tag.Update)
	apiv1.PATCH("/tags/:id/state", tag.Update)
	apiv1.GET("/tags", tag.List)

	apiv1.POST("/articles", article.Create)
	apiv1.DELETE("/articles/:id", article.Delete)
	apiv1.PUT("/articles/:id", article.Update)
	apiv1.PATCH("/articles/:id/state", article.Update)
	apiv1.GET("/articles", article.List)
	apiv1.GET("/articles/:id", article.Get)

	return r
}
