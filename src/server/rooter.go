package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"izanagi-portfolio-site.com/database"
	"izanagi-portfolio-site.com/model"
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

		router.GET("/create", func(c *gin.Context) {
			c.HTML(200, "create.html", gin.H{})
		})

		// ミッドルネーム作成
		router.POST("/create", func(c *gin.Context) {

			var form model.CreatedMiddleNames

			if err := c.Bind(&form); err != nil {
				c.HTML(http.StatusBadRequest, "create.html", gin.H{"err": err})
				c.Abort()
			} else {
				mr := database.DBGetRandomMrData().Mr
				lName := c.PostForm("lname")
				surName := database.DBGetRandomSNData().SurName
				commonName := database.DBGetRandomCNData().CommonName
				fName := c.PostForm("fname")

				session := sessions.Default(c)
				session.Set("mr", mr)
				session.Set("lName", lName)
				session.Set("surName", surName)
				session.Set("commonName", commonName)
				session.Set("fName", fName)
				session.Save()

				// database.DbMiddleNameInsert(mr, lName, surName, commonName, fName)
				c.Redirect(302, "/app/middle_name/result")
			}
		})

		router.GET("/result", func(c *gin.Context) {
			session := sessions.Default(c)
			mr := fmt.Sprintf("%v", session.Get("mr"))
			lName := fmt.Sprintf("%v", session.Get("lName"))
			surName := fmt.Sprintf("%v", session.Get("surName"))
			commonName := fmt.Sprintf("%v", session.Get("commonName"))
			fName := fmt.Sprintf("%v", session.Get("fName"))
			c.HTML(200, "result.html", gin.H{
				"Mr":         mr,
				"LName":      lName,
				"SurName":    surName,
				"CommonName": commonName,
				"FName":      fName,
			})
		})

	}

	routerGlobal.Run(":8001")

	return routerGlobal
}
