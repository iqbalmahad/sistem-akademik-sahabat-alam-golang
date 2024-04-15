package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
// 	"gorm.io/gorm"
// )

// // BaseHandler adalah handler dasar untuk entitas
// type BaseHandler struct {
// 	DB        *gorm.DB
// 	Validator *validator.Validate
// }

// // NewBaseHandler membuat instance baru dari BaseHandler
// func NewBaseHandler(db *gorm.DB, validator *validator.Validate) *BaseHandler {
// 	return &BaseHandler{
// 		DB:        db,
// 		Validator: validator,
// 	}
// }

// // CreateUserHandler menangani pembuatan entitas baru
// func (h *BaseHandler) CreateUserHandler(c *gin.Context) {
// 	var entity models.Entity // Ganti Entity dengan entitas yang sesuai
// 	if err := c.ShouldBindJSON(&entity); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.Validator.Struct(entity); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.DB.Create(&entity).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entity"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, entity)
// }

// // GetEntityHandler menangani pengambilan data entitas berdasarkan ID
// func (h *BaseHandler) GetEntityHandler(c *gin.Context) {
// 	var entity models.Entity // Ganti Entity dengan entitas yang sesuai
// 	entityID := c.Param("id")

// 	if err := h.DB.First(&entity, entityID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, entity)
// }

// // UpdateEntityHandler menangani pembaruan data entitas berdasarkan ID
// func (h *BaseHandler) UpdateEntityHandler(c *gin.Context) {
// 	var entity models.Entity // Ganti Entity dengan entitas yang sesuai
// 	entityID := c.Param("id")

// 	if err := h.DB.First(&entity, entityID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&entity); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.Validator.Struct(entity); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.DB.Save(&entity).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update entity"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, entity)
// }

// // DeleteEntityHandler menangani penghapusan data entitas berdasarkan ID
// func (h *BaseHandler) DeleteEntityHandler(c *gin.Context) {
// 	var entity models.Entity // Ganti Entity dengan entitas yang sesuai
// 	entityID := c.Param("id")

// 	if err := h.DB.First(&entity, entityID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
// 		return
// 	}

// 	if err := h.DB.Delete(&entity).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entity"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Entity deleted successfully"})
// }
