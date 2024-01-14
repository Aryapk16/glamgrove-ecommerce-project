package middleware

import (
	"errors"
	"fmt"
	"glamgrove/pkg/auth"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser(ctx *gin.Context) {
	authHelper(ctx, "user-auth")
}

// user Admin
func AuthenticateAdmin(ctx *gin.Context) {
	authHelper(ctx, "admin-auth")
}

// helper to validate cookie and expiry
func authHelper(ctx *gin.Context, authname string) {

	// get cookie for user or admin with name
	tokenString, err := ctx.Cookie(authname)

	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Error while fetching cookie",
		})
		return
	}

	// auth function validate the token and return claims with error
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login token not valid",
		})
		return
	}

	// check the cliams expire time
	if time.Now().Unix() > claims.ExpiresAt {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "User Need Re-Login time expired",
		})
		return
	}
	// claim the" userId and set it on context
	ctx.Set("userId", fmt.Sprint(claims.Id))

}

//Get user id from JWT Token

func GetId(tokenString string) (uint, error) {
	// Validate the JWT token and retrieve the claims
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	// Extract and parse the user ID from the claims
	userID, err := strconv.ParseUint(claims.Id, 10, 32)
	if err != nil {
		return 0, errors.New("failed to parse user ID from JWT claims")
	}

	return uint(userID), nil
}

// to fetch phone number from jwt
// func GetPhn(ctx *gin.Context authname string) (string, error) {
// 	cookie, err := c.Request.Cookie(authname)
// 	if err != nil {
// 		return " ", errors.New("can't find cookie")
// 	}

// 	tokenString := cookie.Value
// 	claims, err := auth.Validatetoken(tokenString)
// 	if err != nil {
// 		return " ", errors.New("can't validate cookie")
// 	}

// 	phn, ok := claims["sub"].(string)
// 	if !ok {
// 		return " ", errors.New("can't find phn")
// 	}
// 	return phn, nil
// }
