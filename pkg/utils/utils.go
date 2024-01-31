package utils

import (
	"strconv"
	"time"

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

func StringToTime(date string) (time.Time, error) {
	layout := "2006-01-02"

	// Parse the string date using the specified layout
	returnDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	// Return the parsed time
	return returnDate, nil
}
