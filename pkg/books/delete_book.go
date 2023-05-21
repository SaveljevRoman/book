package books

import (
	"book/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Book{}

	if result := h.DB.First(&book, id); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
