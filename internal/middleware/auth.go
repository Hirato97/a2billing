package middleware

import (
	authUtil "a2billing-go-api/common/auth"
	"a2billing-go-api/common/log"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
)

const (
	secretToken = "1j9F5^I0Lr10n*H0Mp2*P^kK@mvv4PR^"
)

var cacheObj libcache.Cache
var strategy union.Union
var tokenStrategy auth.Strategy

type GoAuthMiddleware struct {
	GoAuth authUtil.GoAuth
}

func SetupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 10)
	cacheObj.RegisterOnExpired(func(key, _ interface{}) {
		cacheObj.Peek(key)
	})
	// basicStrategy := basic.NewCached(validateBasicAuth, cacheObj)
	tokenStrategy = token.New(validateTokenAuth, cacheObj)
	strategy = union.New(tokenStrategy)
	// strategy = union.New(tokenStrategy, basicStrategy)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("AuthMiddleware", "AuthMiddleware", "Executing Auth Middleware")
		_, user, err := strategy.AuthenticateRequest(c.Request)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": http.StatusText(http.StatusUnauthorized),
				},
			)
			c.Abort()
			return
		}
		c.Set("user", user)
		log.Info("AuthMiddleware", "Authenticated", user.GetUserName())
	}
}

// func validateBasicAuth(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
// 	userDomain := strings.Split(username, "@")
// 	if len(userDomain) != 2 {
// 		log.Info("AuthMiddleware", "validateBasicAuth", "missing @")
// 		return nil, errors.New("invalid credentials")
// 	}
// 	data, _ := res.(map[string]interface{})
// 	claims, _ := data["data"].(map[string]interface{})
// 	extension := make(map[string][]string)
// 	extension["domain"] = []string{claims["domain_uuid"].(string)}
// 	extension["level"] = []string{claims["level"].(string)}
// 	user := auth.NewDefaultUser(claims["user_uuid"].(string), claims["user_uuid"].(string), nil, extension)
// 	return user, nil
// }

func validateTokenAuth(ctx context.Context, r *http.Request, tokenString string) (auth.Info, time.Time, error) {
	if tokenString == secretToken {
		extension := make(map[string][]string)
		//Just hard code - this param is not thing
		extension["domain"] = []string{"2273f762-7ae6-4a0e-a09d-6d5a3c961a50"}
		extension["level"] = []string{"superadmin"}
		user := auth.NewDefaultUser("portal", "2273f762-7ae6-4a0e-a09d-6d5a3c961a50", nil, extension)
		return user, time.Now(), nil
	}
	client, err := authUtil.GoAuthClient.CheckTokenInRedis(tokenString)
	if err != nil {
		return nil, time.Time{}, err
	}
	token, err := jwt.Parse(client.JWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, time.Time{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		extension := make(map[string][]string)
		user := auth.NewDefaultUser(claims["id"].(string), claims["id"].(string), nil, extension)
		return user, time.Now(), nil
	}
	return nil, time.Time{}, errors.New("invalid token")
}

func GetUserId(c *gin.Context) (interface{}, bool) {
	user, isExist := c.Get("user")
	return user.(auth.Info).GetID(), isExist
}

func GetUserLevel(c *gin.Context) (interface{}, bool) {
	user, isExist := c.Get("user")
	extension := user.(auth.Info).GetExtensions()
	return extension["level"][0], isExist
}

func GetUserDomain(c *gin.Context) (interface{}, bool) {
	user, isExist := c.Get("user")
	extension := user.(auth.Info).GetExtensions()
	return extension["domain"][0], isExist
}

func GetUserName(c *gin.Context) (interface{}, bool) {
	user, isExist := c.Get("user")
	return user.(auth.Info).GetUserName(), isExist
}
