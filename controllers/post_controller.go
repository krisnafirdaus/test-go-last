package controllers

import (
	"net/http"
	"prakerja8/configs"
	"prakerja8/models"

	"github.com/labstack/echo/v4"
)

func GetPostsController(c echo.Context) error {
	var posts []models.Post
	result := configs.DB.Find(&posts)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch posts",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    posts,
	})
}

func CreatePostController(c echo.Context) error {
	var post models.Post
	c.Bind(&post)

	result := configs.DB.Create(&post)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to create post",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Post created successfully",
		Data:    post,
	})
}
