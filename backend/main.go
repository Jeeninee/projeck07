package main

import (
	"github.com/Jeeninee/project07/controller"
	"github.com/Jeeninee/project07/entity"
	"github.com/Jeeninee/project07/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Student Routes
			protected.GET("/students", controller.ListStudents)
			protected.GET("/students/:id", controller.ListStudent)
			protected.GET("/student/:id", controller.GetStudent)
			protected.PATCH("/students", controller.UpdateStudent)
			protected.DELETE("/students/:id", controller.DeleteStudent)

			// Registrar Routes
			protected.GET("/registrars", controller.ListRegistrar)
			protected.GET("/registrar/:id", controller.GetRegistrar)
			protected.POST("/registrar", controller.CreateRegistrar)
			protected.PATCH("/registrars", controller.UpdateRegistrar)
			protected.DELETE("/registrars/:id", controller.DeleteRegistrar)

			// Course Routes
			protected.GET("/courses", controller.ListCourses)
			protected.GET("/course/:id", controller.GetCourse)
			protected.POST("/courses", controller.CreateCourse)
			protected.PATCH("/courses", controller.UpdateCourse)
			protected.DELETE("/courses/:id", controller.DeleteCourse)
		}
	}

	// Student Routes
	r.POST("/students", controller.CreateStudent)

	// Authentication Routes
	r.POST("/student/login", controller.LoginStudent)
	r.POST("/registrar/login", controller.LoginRegistrar)

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