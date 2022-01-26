package controller

import (
	"net/http"

	"github.com/Jeeninee/project07/entity"
	"github.com/gin-gonic/gin"
)

// POST /Course
func CreateCourse(c *gin.Context) {
	var Course entity.Course
	if err := c.ShouldBindJSON(&Course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Course})
}

// GET /Course/:id
func GetCourse(c *gin.Context) {
	var Course entity.Course
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM course WHERE id = ?", id).Scan(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Course})
}

// GET /Courses
func ListCourse(c *gin.Context) {
	var Course []entity.Course
	if err := entity.DB().Raw("SELECT * FROM course").Scan(&Courses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Course})
}

// DELETE /Course/:id
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM course WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Course
func UpdateCourse(c *gin.Context) {
	var Course entity.Course
	if err := c.ShouldBindJSON(&Course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Course.ID).First(&Course); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	if err := entity.DB().Save(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Course})
}