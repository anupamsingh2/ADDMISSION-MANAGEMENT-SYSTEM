package controllers

import (
	"ADMISSION-MANAGEMENT-SYSTEM/config"
	"ADMISSION-MANAGEMENT-SYSTEM/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	student.Password = string(hashedPassword)

	config.DB.Create(&student)

	// Return student data without password
	c.JSON(201, gin.H{
		"id":    student.ID,
		"name":  student.Name,
		"email": student.Email,
		"phone": student.Phone,
	})
}

func GetStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(200, students)
}

func LoginStudent(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	if err := config.DB.Where("email = ?", input.Email).First(&student).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
