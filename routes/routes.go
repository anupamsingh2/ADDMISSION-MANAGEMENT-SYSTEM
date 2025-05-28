package routes

import (
	"ADMISSION-MANAGEMENT-SYSTEM/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	student := r.Group("/students")
	{
		student.POST("/", controllers.CreateStudent)
		student.GET("/", controllers.GetStudents)
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
