package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "SUCCESS",
		"data":   data,
	})
}

func SendSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "SUCCESS",
	})
}

func SendBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":   400,
		"status": "BAD REQUEST",
		"error":  err.Error(),
	})
	c.Abort()
}

func SendInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":   500,
		"status": "INTERNAL SERVER ERROR",
		"error":  err.Error(),
	})
	c.Abort()
}

func SendUnauthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":   401,
		"status": "UNAUTHORIZED",
		"error":  err.Error(),
	})
	c.Abort()
}

func SendNotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":   404,
		"status": "NOT FOUND",
		"error":  err.Error(),
	})
	c.Abort()
}

func SendForbidden(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, gin.H{
		"code":   403,
		"status": "FORBIDDEN",
		"error":  err.Error(),
	})
	c.Abort()
}
