package controllers

import (
	"ADMISSION-MANAGEMENT-SYSTEM/config"
	"ADMISSION-MANAGEMENT-SYSTEM/models"

	"github.com/gin-gonic/gin"
)

func EnrollStudent(c *gin.Context) {
	var enrollment models.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&enrollment)
	c.JSON(201, enrollment)
}

func GetEnrollments(c *gin.Context) {
	var enrollments []models.Enrollment
	config.DB.Find(&enrollments)
	c.JSON(200, enrollments)
}
