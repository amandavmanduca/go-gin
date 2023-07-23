package routes

import (
	"github.com/amandavmanduca/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.HelloWorld)
	r.GET("/alunos", controllers.Get)
	r.GET("/alunos/:id", controllers.GetById)
	r.POST("/alunos", controllers.CreateOne)
	r.DELETE("/alunos/:id", controllers.DeleteOne)
	r.PATCH("/alunos/:id", controllers.UpdateOne)
	r.GET("/alunos/cpf/:cpf", controllers.GetByCPF)
	r.GET("/index", controllers.Index)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
