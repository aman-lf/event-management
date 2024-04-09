package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CurrentUserKey = "userID"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("UserId")
		if userId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx := context.WithValue(c.Request.Context(), CurrentUserKey, userId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GetCurrentUserIDFromContext(ctx context.Context) *int {
	userIdStr, _ := ctx.Value(CurrentUserKey).(string)
	userId, _ := strconv.Atoi(userIdStr)
	return &userId
}
