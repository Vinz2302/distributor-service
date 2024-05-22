package middlewares

import (

	// "github.com/getsentry/sentry-go"
	// sentrygin "github.com/getsentry/sentry-go/gin"

	"bytes"
	"distributor-service/app/config"

	// roleModel "ethical-be/modules/v1/utilities/role/model"
	// res "ethical-be/pkg/api-response"
	// "ethical-be/pkg/discord"
	// "ethical-be/pkg/helpers"
	// jwt "ethical-be/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var (
	env, _              = config.Init()
	authorizationHeader = "Authorization"
	apiKeyHeader        = "X-API-key"
	cronExecutedHeader  = "X-Appengine-Cron"
	valName             = "FIREBASE_ID_TOKEN"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Auth-User-Id, X-Transaction-Id, baggage, sentry-trace")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, HEAD, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

/*
bodyLogWriter for return all response body http request as logging
*/
type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware(ctx *gin.Context) {
	blw := &BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw
	ctx.Next()
	// Store the captured response payload in the context
	ctx.Set("capturedPayload", blw.Body.String())
}
