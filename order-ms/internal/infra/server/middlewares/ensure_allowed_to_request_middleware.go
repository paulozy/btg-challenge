package middlewares

import (
	"encoding/base64"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/paulozy/btg-challenge/order-ms/configs"
)

func EnsureAllowedToRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs, _ := configs.LoadConfig("../../../")
		XOrderMicrosserviceHeader := c.GetHeader("X-Order-Microsservice")

		if XOrderMicrosserviceHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		str := XOrderMicrosserviceHeader
		dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
		n, err := base64.StdEncoding.Decode(dst, []byte(str))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		dst = dst[:n]

		isEqual := reflect.DeepEqual(dst, []byte(configs.HeaderKey))

		if !isEqual {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
