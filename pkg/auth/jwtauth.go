package auth

import (
	"errors"
	"fmt"
	"glamgrove/pkg/config"
	"glamgrove/pkg/domain"
	"glamgrove/pkg/utils/request"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtCookieSetup(c *gin.Context, name string, userId uint) bool {
	//time = 10 mins
	cookieTime := time.Now().Add(10 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{

		Id:        fmt.Sprint(userId),
		ExpiresAt: cookieTime,
	})

	// Generate signed JWT token using env var of secret key
	if tokenString, err := token.SignedString([]byte(config.GetJWTConfig())); err == nil {

		// Set cookie with signed string if no error time = 10 hours
		c.SetCookie(name, tokenString, 10*3600, "", "", false, true)

		fmt.Println("JWT sign & set Cookie successful")
		return true
	}
	fmt.Println("Failed JWT setup")
	return false

}

func ValidateToken(tokenString string) (jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.GetJWTConfig()), nil
		},
	)
	if err != nil || !token.Valid {
		fmt.Println("not valid token")
		return jwt.StandardClaims{}, errors.New("not valid token")
	}

	// then parse the token to claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		fmt.Println("Can't parse the claims")
		return jwt.StandardClaims{}, errors.New("Can't parse the claims")
	}

	return *claims, nil
}

func GenerateTokenForOtp(val domain.User) (string, error) {
	var expiryTime = time.Now().Add(10 * time.Minute).Unix()
	claims := request.OtpCookieStruct{
		FirstName: val.FirstName,
		LastName:  val.LastName,
		Email:     val.Email,
		Phone:     val.Phone,
		UserName:  val.UserName,
		Password:  val.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetJWTConfig()))

	return token, err

}

func ValidateOtpTokens(signedtoken string) (request.OtpCookieStruct, error) {
	token, err := jwt.ParseWithClaims(
		signedtoken, &request.OtpCookieStruct{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTConfig()), nil
		})

	if err != nil {

		return request.OtpCookieStruct{}, err
	}

	claim, _ := token.Claims.(*request.OtpCookieStruct)

	return *claim, nil
}

//sub as phone number
func GenerateJWTPhn(phn string) (map[string]string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": phn,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetJWTConfig()))

	if err != nil {
		return nil, errors.New("JWT token generating is failed")
	}
	return map[string]string{"accessToken": tokenString}, nil
}
