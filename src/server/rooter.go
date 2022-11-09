package server

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"izanagi-portfolio-site.com/database"
)

func NewRouter() *gin.Engine {
	// リリースモード
	// gin.SetMode(gin.ReleaseMode)

	routerGlobal := gin.Default()

	routerGlobal.LoadHTMLGlob("views/*.html")

	router := routerGlobal.Group("app/middle_name")
	{
		router.Static("/store", "./store")
		router.Static("/assets", "./assets")

		store := cookie.NewStore([]byte("secret"))
		router.Use(sessions.Sessions("mysession", store))

		router.GET("/", func(c *gin.Context) {
			c.HTML(200, "home.html", gin.H{})
		})

		//一覧
		router.GET("/history", func(c *gin.Context) {
			middleNames := database.DbGetCreatedMiddleNames()

			c.HTML(http.StatusOK, "history.html", gin.H{
				"name":   "test",
				"middle": middleNames,
			})
		})

		router.GET("/create", func(c *gin.Context) {
			c.HTML(200, "create.html", gin.H{})
		})

		// ミッドルネーム作成
		// router.POST("/create", func(c *gin.Context) {
		// 	session := sessions.Default(c)
		// 	userId := fmt.Sprintf("%v", session.Get("Uuid"))

		// 	var form model.CreatedMiddleNames

		// 	if err := c.Bind(&form); err != nil {
		// 		middleNames := database.DbGetCreatedMiddleNames(userId)
		// 		c.HTML(http.StatusBadRequest, "create.html", gin.H{"middleNames": middleNames, "err": err})
		// 		c.Abort()
		// 	} else {
		// 		mr := database.DBGetRandomMrData().Mr
		// 		lName := c.PostForm("lname")
		// 		surName := database.DBGetRandomSNData().SurName
		// 		commonName := database.DBGetRandomCNData().CommonName
		// 		fName := c.PostForm("fname")

		// 		database.DbMiddleNameInsert(mr, lName, surName, commonName, fName, userId)
		// 		c.Redirect(302, "/app/middle_name/user/result")
		// 	}
		// })

		// router.GET("/result", func(c *gin.Context) {
		// 	session := sessions.Default(c)
		// 	userId := fmt.Sprintf("%v", session.Get("Uuid"))
		// 	middleName := database.DbMiddleNameLastFind(userId)
		// 	c.HTML(200, "result.html", gin.H{
		// 		"middleName": middleName,
		// 	})
		// })

	}

	routerGlobal.Run(":8001")

	return routerGlobal
}
