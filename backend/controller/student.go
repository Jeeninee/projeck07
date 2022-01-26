package controller

import (
	"net/http"

	"github.com/Jeeninee/project07/entity"
	"github.com/gin-gonic/gin"
)

// POST /Student
func CreateStudent(c *gin.Context) {
	var Student entity.Student
	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Student})
}

// GET /Student/:id
func GetStudent(c *gin.Context) {
	var Student entity.Student
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM student WHERE id = ?", id).Scan(&Student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Student})
}

// GET /Students
func ListStudent(c *gin.Context) {
	var Student []entity.Student
	if err := entity.DB().Raw("SELECT * FROM student").Scan(&Students).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Student})
}

// DELETE /Student/:id
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM student WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Student
func UpdateStudent(c *gin.Context) {
	var Student entity.Student
	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Student.ID).First(&Student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	if err := entity.DB().Save(&Student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Student})
}