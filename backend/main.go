package main

import (
	"github.com/Jeeninee/project07/controller"
	"github.com/Jeeninee/project07/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	

			// Semester Routes
			r.GET("/semesters", controller.ListSemesters)
			r.GET("/semester/:id", controller.GetSemester)
			r.POST("/semesters", controller.CreateSemester)
			r.PATCH("/semesters", controller.UpdateSemester)
			r.DELETE("/semesters/:id", controller.DeleteSemester)

			// Grades Routes
			r.GET("/Gradess", controller.ListGradess)
			r.GET("/Grades/:id", controller.GetGrades)
			r.POST("/Gradess", controller.CreateGrades)
			r.PATCH("/Gradess", controller.UpdateGrades)
			r.DELETE("/Gradess/:id", controller.DeleteGrades)			
		

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
