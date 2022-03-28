package main

import (
	"net/http"
	"strings"
	"triadmoko-be-golang/auth"
	"triadmoko-be-golang/config"
	"triadmoko-be-golang/handler"
	"triadmoko-be-golang/helper"
	"triadmoko-be-golang/repository"
	"triadmoko-be-golang/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	// load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	db := config.Database()

	userAuthService := auth.NewService()

	repository := repository.NewRepositoryUser(db)
	service := service.NewServiceUser(repository)
	handler := handler.NewHandlerUser(service, userAuthService)

	router := gin.Default()
	user := router.Group("/api/v1/user")
	faskes := router.Group("/api/v1/faskes")
	nakes := router.Group("/api/v1/nakes")
	user.POST("/register", handler.RegisterUser)
	user.POST("/login", handler.Login)

	faskes.GET("/", authMiddleware(userAuthService, service), handler.FindAllFaskes)
	faskes.POST("/create", handler.CreateFaskes)

	nakes.POST("/create", handler.CreateNakes)
	nakes.PUT("/update/:id", authMiddleware(userAuthService, service), handler.UpdateNakes)
	nakes.GET("/", handler.FindAllNakes)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func authMiddleware(userAuthService auth.Service, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponseApi("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := auth.NewService().ValidateToken(tokenString)
		if err != nil {
			response := helper.ResponseApi("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, oke := token.Claims.(jwt.MapClaims)
		if !oke || !token.Valid {
			response := helper.ResponseApi("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := service.GetUserByID(userID)
		if err != nil {
			response := helper.ResponseApi("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currenUser", user)
	}
}
