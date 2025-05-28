package routes

import (
	"ADMISSION-MANAGEMENT-SYSTEM/controllers"
	"ADMISSION-MANAGEMENT-SYSTEM/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	student := r.Group("/students")
	{
		student.POST("/", controllers.CreateStudent)
		student.GET("/", controllers.GetStudents)
		student.POST("/login", controllers.LoginStudent)

		// Protected routes under /students that require JWT auth
		protected := student.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		protected.GET("/profile", controllers.StudentProfile) // example protected route
	}

	course := r.Group("/courses")
	{
		course.POST("/", controllers.CreateCourse)
		course.GET("/", controllers.GetCourses)
	}

	enrollment := r.Group("/enrollments")
	{
		enrollment.POST("/", controllers.EnrollStudent)
		enrollment.GET("/", controllers.GetEnrollments)
	}
}
