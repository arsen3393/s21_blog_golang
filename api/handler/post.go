package handler

import (
	"Go_Day06/api/middleware"
	"Go_Day06/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPostRoutes(router *gin.Engine, postModel *models.PostModel) {
	router.GET("/", GetPosts(postModel))
	protected := router.Group("/post")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/", CreatePost(postModel))
	}
}

type GetPostRequest struct {
	Page int `form:"page" binding:"required"`
}

type PostResponse struct {
	Posts     []models.Post `json:"posts"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	HasNext   bool          `json:"hasNext"`
	HasPrev   bool          `json:"hasPrev"`
}

// GetPosts получает список постов
// @Summary Получение постов
// @Description Возвращает список постов с пагинацией
// @Tags posts
// @Accept json
// @Produce json
// @Param page query int true "Номер страницы"
// @Success 200 {object} PostResponse
// @Failure 400 {object} map[string]interface{} "Ошибка запроса"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router / [get]
func GetPosts(postModel *models.PostModel) gin.HandlerFunc {
	const op = "GetPostsHandler"
	return func(c *gin.Context) {
		var req GetPostRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			fmt.Println(op+" : ", err)
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		posts, total, err := postModel.GetAllPosts(req.Page)
		if err != nil {
			fmt.Println(op+" : ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err})
			return
		}
		pagSize := 3
		totalPages := (total + pagSize - 1) / pagSize
		hasNext := req.Page < totalPages
		hasPrev := req.Page > 1
		response := PostResponse{
			Posts:     posts,
			PageSize:  pagSize,
			TotalPage: totalPages,
			HasNext:   hasNext,
			HasPrev:   hasPrev,
		}
		c.JSON(http.StatusOK, response)
	}
}

type PutPostRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"text" binding:"required"`
}

type PutPostResponse struct {
	Response string `json:"response"`
}

// CreatePost добавляет новый пост
// @Summary Создание поста
// @Description Создаёт новый пост в блоге (нужен токен)
// @Tags posts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param post body PutPostRequest true "Данные поста"
// @Success 200 {object} PutPostResponse
// @Failure 400 {object} map[string]interface{} "Ошибка запроса"
// @Failure 401 {object} map[string]interface{} "Неавторизован"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /post/ [post]
func CreatePost(pm *models.PostModel) gin.HandlerFunc {
	const op = "PutPostHandler"
	return func(c *gin.Context) {
		var req PutPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(op+" : ", err)
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		err := pm.CreatePost(req.Title, req.Body)
		if err != nil {
			fmt.Println(op+" : ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err})
			return
		}
		res := PutPostResponse{
			Response: "ok",
		}
		c.JSON(http.StatusOK, res)
	}
}
