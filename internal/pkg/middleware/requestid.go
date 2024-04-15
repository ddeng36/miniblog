package middleware

import (
	"github.com/ddeng36/miniblog/internal/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		requestID := c.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 RequestID 保存在 gin.Context 中，方便后边程序使用
		c.Set(known.XRequestIDKey, requestID)

		// 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID`
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		c.Next()
	}
}