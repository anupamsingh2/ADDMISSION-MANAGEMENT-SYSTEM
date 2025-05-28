package controllers

import (
	"ADMISSION-MANAGEMENT-SYSTEM/config"
	"ADMISSION-MANAGEMENT-SYSTEM/models"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&student)
	c.JSON(201, student)
}

func GetStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(200, students)
}
