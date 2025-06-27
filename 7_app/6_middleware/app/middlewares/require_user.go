package middlewares

import (
	"errors"
	"fmt"
	"learning/app/util/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func RequireUser(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			log.Info().Err(err).Msg("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, AbortError(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			log.Info().Err(err).Msg("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, AbortError(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			log.Info().Err(err).Msg("unsupported authorization type")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, AbortError(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			log.Info().Err(err).Msg("error verifying token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, AbortError(err))
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
