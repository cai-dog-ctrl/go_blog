package routers

import (
	"blog/global"
	"blog/internal/middleware"
	"blog/internal/routers/api"
	v1 "blog/internal/routers/api/v1"
	"blog/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations()) //注册Translation中间件
	r.Use(gin.Logger())              //Logger 实例一个 Logger 中间件，它将日志写入gin
	r.Use(gin.Recovery())            //Recovery 返回一个中间件，它可以从任何恐慌中恢复，如果有，则写入 500
	r.Use(middleware.Tracer())       //链路追踪
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url)) //注册一个swagger路由访问
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile) //上传文件
	r.POST("/auth", api.GetAuth)              //校验JWT
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1") //创建路由组
	apiv1.Use(middleware.JWT()) //apiv1 路由组全部加入JWT中间件
	{
		apiv1.POST("/tags", tag.Create)            //新增标签
		apiv1.DELETE("/tags/:id", tag.Delete)      //删除标签
		apiv1.PUT("/tags/:id", tag.Update)         //更新指定标签
		apiv1.PATCH("/tags/:id/state", tag.Update) //更新标签列表
		apiv1.GET("/tags", tag.List)               //获取指定标签

		apiv1.POST("/articles", article.Create)            //新增文章
		apiv1.DELETE("/articles/:id", article.Delete)      //删除文章
		apiv1.PUT("/articles/:id", article.Update)         //更新指定文章
		apiv1.PATCH("/articles/:id/state", article.Update) //
		apiv1.GET("/articles/:id", article.Get)            //获取指定文章
		apiv1.GET("/articles", article.List)               //获取文章列表
	}

	return r
}
