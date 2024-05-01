package auth

import (
	"errors"
	"fmt"
	"net/http"
	"projectsphere/cats-social/pkg/protocol/msg"
	"projectsphere/cats-social/pkg/utils/config"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var key = config.Get().Auth.SecretKey
var accessTokenExpiredTime = config.Get().Auth.AccessTokenExpiredTime
var refreshTokenExpiredTime = config.Get().Auth.RefreshTokenExpiredTime

func GenerateToken(userId uint32, tokenType string) (string, error) {
	claims := jwt.MapClaims{}
	claims["issuer"] = "JWT_issuer"
	claims["userId"] = userId
	claims["issuedAt"] = time.Now().Unix()

	if tokenType == "ACCESS_TOKEN" {
		accessTokenExpiredDuration, err := time.ParseDuration(accessTokenExpiredTime)
		if err != nil {
			return "", msg.InternalServerError(err.Error())
		}
		claims["exp"] = time.Now().Add(accessTokenExpiredDuration).Unix()
		claims["tokenType"] = "ACCESS_TOKEN"
	} else if tokenType == "REFRESH_TOKEN" {
		refreshTokenExpiredDuration, err := time.ParseDuration(refreshTokenExpiredTime)
		if err != nil {
			return "", msg.InternalServerError(err.Error())
		}
		claims["exp"] = time.Now().Add(refreshTokenExpiredDuration).Unix()
		claims["tokenType"] = "REFRESH_TOKEN"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

// func GenerateAccessTokenByRefreshToken(c *gin.Context) (string, error) {
// 	tokenString := ExtractToken(c)
// 	if tokenString == "" {
// 		return "", errors.New(msg.ErrTokenNotExist)
// 	}

// 	tokenData, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(key), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	var role int
// 	var userId uint
// 	var currTokenType string
// 	var expireAt int64

// 	claims, ok := tokenData.Claims.(jwt.MapClaims)
// 	if ok && tokenData.Valid {
// 		role = int(claims["roleId"].(float64))
// 		userId = uint(claims["userId"].(float64))
// 		currTokenType = claims["tokenType"].(string)
// 		expireAt = int64(claims["exp"].(float64))

// 	}

// 	if currTokenType != "REFRESH_TOKEN" {
// 		return "", fmt.Errorf(msg.ErrInvalidTokenType)
// 	}

// 	if role != 1 && role != 2 {
// 		return "", errors.New(msg.ErrUserRoleNotExist)
// 	}

// 	isTimeValid := checkTokenTimeValid(expireAt)
// 	if !isTimeValid {
// 		return "", fmt.Errorf(msg.ErrTokenAlreadyExpired)
// 	}

// 	res, err := GenerateToken(userId, "ACCESS_TOKEN")
// 	if err != nil {
// 		return "", err
// 	}
// 	return res, nil
// }

func TokenValid(c *gin.Context, isAdminRole bool, isUserRole bool) error {
	tokenString := ExtractToken(c)
	if tokenString == "" {
		return errors.New(msg.ErrTokenNotExist)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil && err.Error() == "signature is invalid" {
		return errors.New(msg.ErrTokenAlreadyExpired)
	} else if err != nil {
		return err
	}

	var role int
	var currTokenType string
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		role = int(claims["roleId"].(float64))
		currTokenType = claims["tokenType"].(string)
	}

	if role != 1 && role != 2 {
		return errors.New(msg.ErrUserRoleNotExist)
	}

	if isAdminRole && role != 1 && isUserRole {
		return fmt.Errorf(msg.ErrUnauthorizedAction)
	}

	if !isAdminRole && role != 2 && isUserRole {
		return fmt.Errorf(msg.ErrUnauthorizedAction)
	}

	if currTokenType != "ACCESS_TOKEN" {
		return fmt.Errorf(msg.ErrInvalidTokenType)
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractUserID(c *gin.Context, tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func ExtractTokenType(c *gin.Context, tokenString string) (string, error) {
	tokenType, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokenType.Claims.(jwt.MapClaims)
	if ok && tokenType.Valid {
		return claims["tokenType"].(string), nil
	}
	return "", errors.New(msg.ErrInvalidToken)
}

func GetUserId(c *gin.Context) (uint, error) {
	token := ExtractToken(c)
	if token == "" {
		return 0, errors.New(msg.ErrTokenNotFound)
	}
	return ExtractUserID(c, token)
}

func JwtAuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c, false, true)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msg.Unauthorization(err.Error()))
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthSuperUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c, true, true)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msg.Unauthorization(err.Error()))
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthNoRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c, false, false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msg.Unauthorization(err.Error()))
			c.Abort()
			return
		}
		c.Next()
	}
}

func checkTokenTimeValid(timestamp interface{}) bool {
	currTime := time.Now()
	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := tm.Sub(currTime)
		if remainder > 0 {
			return true
		}
	}
	return false
}
