package controllers

import (
	"ADMISSION-MANAGEMENT-SYSTEM/config"
	"ADMISSION-MANAGEMENT-SYSTEM/models"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&course)
	c.JSON(201, course)
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	config.DB.Find(&courses)
	c.JSON(200, courses)
}
