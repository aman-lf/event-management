package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// {"method":"GET","path":"/v1/demo","remoteAddr":"::1","responseTime":"1.2ms","startTime":"2023/12/12 - 12:22:11","statusCode":200}
func JSONLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["statusCode"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["requestTime"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remoteAddr"] = params.ClientIP
			log["responseTime"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}
