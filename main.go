package main

import (
	"github.com/cedy/simdocs/controllers"
	"github.com/cedy/simdocs/models"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()
	r.HTMLRender = createRenderer()
	r.GET("/records/id/:id", controllers.GetRecord)
	r.GET("/records", controllers.GetAllRecords)
	r.POST("/records", controllers.CreateRecord)
	r.GET("/records/search", controllers.GetRecordsSearch)
	r.PATCH("/records/:id", controllers.UpdateRecord)
	r.DELETE("/records/:id", controllers.DeleteRecord)

	r.Run()

}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/Base.html", "templates/Index.html", "templates/Navbar.html")
	return r
}
