package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func StringToUint(str string) (uint, error) {
	val, err := strconv.Atoi(str)
	return uint(val), err
}

func GetUserIdFromContext(ctx *gin.Context) uint {
	userIdStr, exists := ctx.Get("userId")
	if !exists {
		return 0
	}
	userId, _ := strconv.ParseUint(userIdStr.(string), 10, 64)
	return uint(userId)
}
