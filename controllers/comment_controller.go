package controllers

import (
	"net/http"
	"prakerja8/configs"
	"prakerja8/models"

	"github.com/labstack/echo/v4"
)

func GetCommentsByPostController(c echo.Context) error {
	var comments []models.Comment
	postId := c.Param("post_id")
	result := configs.DB.Where("post_id = ?", postId).Find(&comments)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch comments",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    comments,
	})
}

func AddCommentController(c echo.Context) error {
	var comment models.Comment
	c.Bind(&comment)

	result := configs.DB.Create(&comment)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to add comment",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Comment added successfully",
		Data:    comment,
	})
}
