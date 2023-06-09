package books

import (
	"book/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}
