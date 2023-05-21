package books

import (
	"book/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *gin.Context) {
	body := AddBookRequestBody{}

	//Запрос бади
	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
