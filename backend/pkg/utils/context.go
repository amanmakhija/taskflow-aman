package utils

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("user_id")
	if !exists {
		return ""
	}

	return userID.(string)
}
