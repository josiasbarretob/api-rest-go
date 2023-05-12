package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IniciarRotas() {
	// Inicializa o Router utilizando as configurações Default do Gin.
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
