package middleware

import (
	"coinpe/pkg/constants"
	"coinpe/pkg/jwtauth"
	"coinpe/pkg/logger"
	"net/http"
	"strings"

	errorConst "coinpe/pkg/error"

	"github.com/gin-gonic/gin"
)

func parseBearerToken(c *gin.Context) string {
	authHeader := c.GetHeader(constants.AuthorizationHeaderName)
	tokenString, found := strings.CutPrefix(authHeader, "Bearer")
	if found {
		tokenString = strings.TrimSpace(tokenString)
	}
	return tokenString
}

func AccessTokenMiddleware(secretKey []byte, allowPartial bool, allowNoAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			errResponse = errorConst.ErrorResponse{}
		)

		tokenString := parseBearerToken(c)
		if tokenString == "" {
			logger.Error("auth header cannot be empty")
			if allowNoAuth {
				// if no auth is there  and token is not present allow them directly
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse.Generate(
					errorConst.ErrorForbidden,
					errorConst.ErrorText(errorConst.ErrorForbidden),
					errorConst.EmptyInterface,
				))
			}
			return
		}

		token, err := jwtauth.ParseToken(tokenString, secretKey)
		if err != nil {
			logger.Error("error in parsing access token | err: ", err)
			if allowNoAuth {
				// if no auth is there  and token is not present allow them directly
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse.Generate(
					errorConst.ErrorForbidden,
					errorConst.ErrorText(errorConst.ErrorForbidden),
					errorConst.EmptyInterface,
				))
			}
			return
		}

		if token.IsPartial && !allowPartial {
			logger.Error("cannot mix tokentype partial with full auth scoped token")
			c.AbortWithStatusJSON(http.StatusForbidden, errResponse.Generate(
				errorConst.ErrorForbidden,
				errorConst.ErrorText(errorConst.ErrorForbidden),
				errorConst.EmptyInterface,
			))
			return
		}

		if !token.IsPartial && allowPartial {
			logger.Error("cannot mix tokentype partial with full auth scoped token")
			c.AbortWithStatusJSON(http.StatusForbidden, errResponse.Generate(
				errorConst.ErrorForbidden,
				errorConst.ErrorText(errorConst.ErrorForbidden),
				errorConst.EmptyInterface,
			))
			return
		}

		c.Set(constants.AuthorizedAccountUUIDContextKey, token.AccountUUID)
		c.Set(constants.AuthorizedAccountRoleContextKey, token.Role)
		c.Set(constants.IsPartialContextKey, token.IsPartial)
		c.Next()
	}

}
