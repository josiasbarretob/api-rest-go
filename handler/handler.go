package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET message",
	})
}

func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET message",
	})
}

func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET message",
	})
}

func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET message",
	})
}

func ListOpeningsHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET message",
	})
}