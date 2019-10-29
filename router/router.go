package router

import (
	"blog/http/api"
	"blog/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init(e *gin.Engine) {

	e.Use(middleware.Cors())

	// 路由
	g := e.Group("/")
	{
		g.StaticFS("/upload/images", http.Dir("runtime/upload/images"))
		g.GET("/captcha", api.GetCaptcha)
		g.GET("/captcha/:captcha_id", api.GetCaptchaImg)

		g.GET("/ping", api.Ping)

		g.POST("/register", api.UserRegister)
		g.POST("/login", api.UserLogin)

		g.GET("/article/:id", api.ArticleGet)

		g.GET("/articles", api.ArticleGetAll)
		g.GET("/articles-topics", api.ArticlesGetWithTopicID)
		g.GET("/articles-tags", api.ArticlesGetWithTagID)
		g.GET("/topic-articles", api.GetArticlesbyArticleID)

		g.GET("/tags", api.TagsGetAll)
		g.GET("/topics", api.TopicsGetAll)
	}

	u := e.Group("/user")
	u.Use(middleware.JwtAuth())
	{
		u.GET("/info", api.UserMe)
		u.PUT("/info/:id", api.UserInfoModify)

	}

	a := e.Group("/admin")
	{
		a.POST("/login", api.AdminLogin)

		a.Use(middleware.JwtAuth())
		{
			a.POST("/register", api.AdminRegister)
			a.GET("/info", api.AdminMe)
			a.PUT("/info/:id", api.AdminInfoModify)

			a.GET("/admins", api.AdminGetAll)
			a.GET("/users", api.UserGetAll)

			a.PUT("/article/:id", api.ArticleEdit)
			a.PUT("/tag/:id", api.TagEdit)
			a.PUT("/topic/:id", api.TopicEdit)

			a.POST("/topic", api.TopicAdd)
			a.POST("/tag", api.TagAdd)
			a.POST("/article", api.ArticleAdd)
			a.DELETE("/article/:id", api.ArticleDelete)
			a.DELETE("/tag/:id", api.TagDelete)
			a.DELETE("/topic/:id", api.TopicDelete)

			a.DELETE("/user/:id", api.UserDelete)
			a.DELETE("/admin/:id", api.AdminDelete)
			a.POST("/upload", api.UploadImage)
		}
	}

}
