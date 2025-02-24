package handler

import (
	"Go_Day06/models"
	"Go_Day06/pkg/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserRoutes(router *gin.Engine, userModel *models.UserModel) {
	user := router.Group("/user")
	{
		// Получение пользователя по ID
		// @Summary Получить пользователя по ID
		// @Description Получает информацию о пользователе по заданному ID
		// @Tags users
		// @Accept json
		// @Produce json
		// @Param id path int true "ID пользователя"
		// @Success 200 {object} GetUserByIdResponse
		// @Failure 400 {object} map[string]interface{} "Неверный формат запроса"
		// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
		// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
		// @Router /user/{id} [get]
		user.GET(":id", GetUserById(userModel))
	}
	// Авторизация пользователя
	// @Summary Авторизация пользователя
	// @Description Авторизует пользователя по имени и паролю и генерирует JWT токен
	// @Tags users
	// @Accept json
	// @Produce json
	// @Param login body UserLoginRequest true "Данные для входа"
	// @Success 200 {object} UserLoginResponse
	// @Failure 400 {object} map[string]interface{} "Ошибка запроса"
	// @Failure 401 {object} map[string]interface{} "Неверные учетные данные"
	// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
	// @Router /login [post]
	router.POST("/login", Login(userModel))
}

type GetUserByIdRequest struct {
	ID int `uri:"id" binding:"required"` // ID пользователя из пути запроса
}

type GetUserByIdResponse struct {
	Name  string `json:"name"`  // Имя пользователя
	Email string `json:"email"` // Email пользователя
}

// GetUserById получает информацию о пользователе по его ID
// @Summary Получить пользователя по ID
// @Description Возвращает информацию о пользователе по заданному ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} GetUserByIdResponse
// @Failure 400 {object} map[string]interface{} "Неверный формат запроса"
// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /user/{id} [get]
func GetUserById(userModel *models.UserModel) gin.HandlerFunc {
	const op = "get_user_by_id_handler"
	return func(c *gin.Context) {
		var req GetUserByIdRequest
		if err := c.ShouldBindUri(&req); err != nil {
			fmt.Printf("%s : %s\n", op, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userModel.GetUserById(req.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := GetUserByIdResponse{
			Name:  user.Name,
			Email: user.Email,
		}
		c.JSON(http.StatusOK, response)
	}
}

type UserLoginRequest struct {
	Name     string `json:"name" binding:"required"`     // Имя пользователя
	Password string `json:"password" binding:"required"` // Пароль пользователя
}

type UserLoginResponse struct {
	Token string `json:"token"` // JWT токен пользователя
}

// Login авторизует пользователя и генерирует JWT токен
// @Summary Авторизация пользователя
// @Description Авторизует пользователя и генерирует JWT токен
// @Tags users
// @Accept json
// @Produce json
// @Param login body UserLoginRequest true "Данные для входа"
// @Success 200 {object} UserLoginResponse
// @Failure 400 {object} map[string]interface{} "Ошибка запроса"
// @Failure 401 {object} map[string]interface{} "Неверные учетные данные"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /login [post]
func Login(userModel *models.UserModel) gin.HandlerFunc {
	const op = "LoginHandler"
	return func(c *gin.Context) {
		var req UserLoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(op+" : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверяем, существует ли пользователь
		user, err := userModel.GetUserByName(req.Name)
		if err != nil || user.Password != req.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Генерация JWT токена
		token, err := auth.GenerateToken(user.ID, user.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		// Отправка токена в ответе
		c.JSON(http.StatusOK, UserLoginResponse{Token: token})
	}
}
