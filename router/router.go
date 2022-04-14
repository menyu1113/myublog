package router

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "myublog/api/v1"
	"myublog/global/vipers"
	"myublog/middleware"
	"myublog/middleware/ZapLog"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}
func InitRouter() {
	if !vipers.AppDebug {
		gin.SetMode("release")
	}
	r := gin.New()
	//r.HTMLRender = createMyRender()
	r.Use(gin.Recovery())
	r.Use(ZapLog.ZapLogs()) //日志中间件
	if vipers.AllowCrossDomain {
		r.Use(middleware.Cors())
	}

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	//后台管理路由接口
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.QueryUserList)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		auth.PUT("user/changepasswd/:id", v1.ChangeUserPassword)
		//// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add/:id", v1.AddCategory)
		auth.PUT("category/:cateid", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//// 文章模块的路由接口
		auth.GET("article/info/:id", v1.GetArtInfo)
		auth.GET("article/:id", v1.GetArt)
		auth.POST("article/add/:id", v1.AddArticle)
		auth.PUT("article/:userid/:artid", v1.EditArt)
		auth.DELETE("article/:artid", v1.DeleteArt)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
		//// 更新个人详情
		auth.POST("adddetail/:id", v1.AddUserDetail)//为用户加用户详情
		auth.GET("admin/detail/:id", v1.GetDetails)//获取用户详情
		auth.PUT("detail/:id", v1.UpdateDetails)//更改用户详情
		//// 评论模块
		auth.POST("addcomment/:id", v1.AddComment)
		auth.GET("comment/", v1.GetComment)
		auth.DELETE("delcomment/:id", v1.DeleteComment)

	}
	public := r.Group("api/v1")
	{
		public.POST("user/add", v1.UserRegister)
		public.POST("login", v1.Login)
	}
	_ = r.Run(vipers.HttpIp + vipers.HttpPort)

}
