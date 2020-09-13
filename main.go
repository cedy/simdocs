package main

import (
	"os"

	"github.com/cedy/simdocs/controllers"
	"github.com/cedy/simdocs/models"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var user, password, port string

func init() {
	if _, err := os.Stat("docs"); os.IsNotExist(err) {
		err := os.Mkdir("docs", 0755)
		if err != nil {
			panic(err)
		}
	}
	if len(os.Args) >= 3 {
		user = os.Args[1]
		password = os.Args[2]
		if len(os.Args) >= 4 {
			port = ":" + os.Args[3]
		}
	}
}

func main() {
	r := gin.Default()
	auth := r.Group("/")
	if user != "" {
		auth = r.Group("/", gin.BasicAuth(gin.Accounts{
			user: password,
		}))
	}
	models.ConnectDataBase()
	// templates, static and uploaded files
	auth.Static("/css", "./assets/css")
	auth.Static("/js", "./assets/js")
	auth.Static("/docs", "./docs")
	r.HTMLRender = createRenderer()
	// CRUD routes
	auth.GET("/records/id/:id", controllers.GetRecord)
	auth.GET("/records", controllers.GetAllRecords)
	auth.GET("/records/create", controllers.CreateRecordForm)
	auth.POST("/records/create", controllers.CreateRecord)
	auth.GET("/records/search", controllers.GetRecordsSearch)
	auth.GET("/records/edit/:id", controllers.EditRecordForm)
	auth.PUT("/records/edit", controllers.UpdateRecord)
	auth.DELETE("/records/:id", controllers.DeleteRecord)
	auth.DELETE("/files/:id", controllers.DeleteFile)
	r.Run(port)

}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/Base.html", "templates/Index.html", "templates/Navbar.html")
	r.AddFromFiles("create", "templates/Base.html", "templates/Create.html", "templates/Navbar.html")
	r.AddFromFiles("edit", "templates/Base.html", "templates/Edit.html", "templates/Navbar.html")
	r.AddFromFiles("record", "templates/Base.html", "templates/Record.html", "templates/Navbar.html")
	return r
}
