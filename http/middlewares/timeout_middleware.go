package middlewares

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxT, cancel := context.WithTimeout(ctx.Request.Context(), timeout)
		defer cancel()

		// Create a channel to signal when processing is done
		done := make(chan struct{})

		go func() {
			// Run the next middleware/handler
			ctx.Next()
			close(done) // Signal that processing is done
		}()

		select {
		case <-ctxT.Done():
			if errors.Is(ctxT.Err(), context.DeadlineExceeded) {
				ctx.JSON(http.StatusGatewayTimeout, gin.H{
					"error": "request timeout",
				})
				ctx.Abort()
			}
		case <-done: // Handler finished without timeout
			// Do nothing, allow the response to proceed
		}
	}
}
