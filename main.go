package main

import (
	"os"

	"github.com/cedy/simdocs/controllers"
	"github.com/cedy/simdocs/models"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func init() {
	if _, err := os.Stat("docs"); os.IsNotExist(err) {
		err := os.Mkdir("docs", 0755)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	r := gin.Default()
	models.ConnectDataBase()
	// templates, static and uploaded files
	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.Static("/docs", "./docs")
	r.HTMLRender = createRenderer()
	// CRUD routes
	r.GET("/records/id/:id", controllers.GetRecord)
	r.GET("/records", controllers.GetAllRecords)
	r.GET("/records/create", controllers.CreateRecordForm)
	r.POST("/records/create", controllers.CreateRecord)
	r.GET("/records/search", controllers.GetRecordsSearch)
	r.GET("/records/edit/:id", controllers.EditRecordForm)
	r.PUT("/records/edit", controllers.UpdateRecord)
	r.DELETE("/records/:id", controllers.DeleteRecord)
	r.DELETE("/files/:id", controllers.DeleteFile)
	r.Run()

}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/Base.html", "templates/Index.html", "templates/Navbar.html")
	r.AddFromFiles("create", "templates/Base.html", "templates/Create.html", "templates/Navbar.html")
	r.AddFromFiles("edit", "templates/Base.html", "templates/Edit.html", "templates/Navbar.html")
	r.AddFromFiles("record", "templates/Base.html", "templates/Record.html", "templates/Navbar.html")
	return r
}
