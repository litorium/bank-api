package middleware

import (
	"bank-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// check exist token
		h := &authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)

		// check token kosong
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		// check verify token
		token, err := utils.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
