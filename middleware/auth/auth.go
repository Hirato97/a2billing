package auth

import (
	"a2billing-go-api/common/log"
	goauth "a2billing-go-api/middleware/auth/goauth"
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
	GoAuth goauth.GoAuth
}

type GoAuthInfo interface {
	auth.Info
}

type GoAuthUser struct {
	Name       string          `json:"name"`
	ID         string          `json:"id"`
	Groups     []string        `json:"groups"`
	Extensions auth.Extensions `json:"extensions"`
	Level      string          `json:"level"`
	Scopes     []string        `json:"scopes"`
}

func (d *GoAuthUser) GetUserName() string {
	return d.Name
}

func (d *GoAuthUser) SetUserName(name string) {
	d.Name = name
}

func (d *GoAuthUser) GetID() string {
	return d.ID
}

func (d *GoAuthUser) SetID(id string) {
	d.ID = id
}

func GetUser(c *gin.Context) (*GoAuthUser, bool) {
	tmp, isExist := c.Get("user")
	if isExist {
		user, ok := tmp.(*GoAuthUser)
		return user, ok
	} else {
		return nil, false
	}
}

func GetUserId(c *gin.Context) (string, bool) {
	user, ok := GetUser(c)
	if !ok {
		return "", false
	} else {
		return user.ID, true
	}
}

func GetUserLevel(c *gin.Context) (string, bool) {
	user, ok := GetUser(c)
	if !ok {
		return "", false
	} else {
		return user.Level, true
	}
}

func GetUserName(c *gin.Context) (interface{}, bool) {
	user, ok := GetUser(c)
	if !ok {
		return "", false
	} else {
		return user.Name, true
	}
}

func (d *GoAuthUser) GetGroups() []string {
	return d.Groups
}

func (d *GoAuthUser) SetGroups(groups []string) {
	d.Groups = groups
}

func (d *GoAuthUser) GetExtensions() auth.Extensions {
	if d.Extensions == nil {
		d.Extensions = auth.Extensions{}
	}
	return d.Extensions
}

func (d *GoAuthUser) SetExtensions(exts auth.Extensions) {
	d.Extensions = exts
}

func (a *GoAuthUser) SetLevel(level string) {
	a.Level = level
}

func (a *GoAuthUser) GetLevel() string {
	return a.Level
}

func (a *GoAuthUser) SetScopes(scopes []string) {
	a.Scopes = scopes
}

func (a *GoAuthUser) GetScopes() []string {
	return a.Scopes
}

func NewGoAuthUser(name, id string, groups []string, extensions auth.Extensions, level string, scopes []string) GoAuthInfo {
	user := &GoAuthUser{
		Level:  level,
		Scopes: scopes,
	}
	user.Name = name
	user.ID = id
	user.Groups = groups
	user.Extensions = extensions
	return user
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
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, user, err := strategy.AuthenticateRequest(c.Request)
		if err != nil {
			log.Error("AuthMiddleware", "validateBasicAuth", "invalid credentials")
			c.JSON(
				http.StatusUnauthorized,
				map[string]interface{}{
					"error": http.StatusText(http.StatusUnauthorized),
				},
			)
			c.Abort()
			return
		}
		c.Set("user", user)

	}
}

/*
func validateBasicAuth(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
	userDomain := strings.Split(username, "@")
	if len(userDomain) != 2 {
		log.Error("AuthMiddleware", "validateBasicAuth", "missing @")
		return nil, errors.New("invalid credentials")
	}
	var domain string
	username = userDomain[0]
	domain = userDomain[1]
	code, res := service.UserServiceGlobal.Authen(domain, username, password)
	if code != http.StatusOK {
		log.Error("AuthMiddleware", "validateBasicAuth", "invalid credentials")
		return nil, errors.New("invalid credentials")
	}
	data, _ := res.(map[string]interface{})
	claims, _ := data["data"].(map[string]interface{})
	id, _ := claims["user_uuid"].(string)
	name, _ := claims["username"].(string)
	domainId, _ := claims["domain_uuid"].(string)
	domainName, _ := claims["domain_name"].(string)
	level, _ := claims["level"].(string)
	user := NewGoAuthUser(name, id, nil, nil, domainId, domainName, level, nil)
	return user, nil
}
*/

func validateTokenAuth(ctx context.Context, r *http.Request, tokenString string) (auth.Info, time.Time, error) {
	if tokenString == secretToken {
		id := "2273f762-7ae6-4a0e-a09d-6d5a3c961a50"
		name := "portal"
		level := "superadmin"
		user := NewGoAuthUser(name, id, nil, nil, level, nil)
		return user, time.Now(), nil
	}
	client, err := goauth.GoAuthClient.CheckTokenInRedis(tokenString)
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
		id := client.UserId
		name, _ := claims["username"].(string)
		level, _ := claims["level"].(string)
		user := NewGoAuthUser(name, id, nil, nil, level, nil)
		return user, time.Now(), nil
	}
	return nil, time.Time{}, errors.New("invalid token")
}
