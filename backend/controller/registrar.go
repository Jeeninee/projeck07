package controller

import (
	"net/http"

	"github.com/Jeeninee/project07/entity"
	"github.com/gin-gonic/gin"
)

// POST /Registrar
func CreateRegistrar(c *gin.Context) {
	var Registrar entity.Registrar
	if err := c.ShouldBindJSON(&Registrar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Registrar).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Registrar})
}

// GET /Registrar/:id
func GetRegistrar(c *gin.Context) {
	var Registrar entity.Registrar
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM registrar WHERE id = ?", id).Scan(&Registrar).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Registrar})
}

// GET /Registrars
func ListRegistrar(c *gin.Context) {
	var Registrar []entity.Registrar
	if err := entity.DB().Raw("SELECT * FROM registrar").Scan(&Registrars).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Registrar})
}

// DELETE /Registrar/:id
func DeleteRegistrar(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM registrar WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registrar not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Registrar
func UpdateRegistrar(c *gin.Context) {
	var Registrar entity.Registrar
	if err := c.ShouldBindJSON(&Registrar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Registrar.ID).First(&Registrar); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registrar not found"})
		return
	}

	if err := entity.DB().Save(&Registrar).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Registrar})
}