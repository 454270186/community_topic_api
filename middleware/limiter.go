package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenBucket struct {
	Capacity  int64
	Rate      float64
	TokenCnt  float64
	LastToken time.Time
	Mtx       sync.Mutex
}

func (tb *TokenBucket) Allow() bool {
	tb.Mtx.Lock()
	defer tb.Mtx.Unlock()

	now := time.Now()
	tb.TokenCnt = tb.TokenCnt + tb.Rate * now.Sub(tb.LastToken).Seconds()

	if tb.TokenCnt >= 1 {
		tb.TokenCnt--
		tb.LastToken = time.Now()
		return true
	} else {
		return false
	}
}

func LimitMid(maxConn int64) gin.HandlerFunc {
	tb := TokenBucket{
		Capacity: maxConn,
		Rate: 1.0,
		TokenCnt: 0,
		LastToken: time.Now(),
	}
	
	return func(ctx *gin.Context) {
		if !tb.Allow() {
			ctx.AbortWithStatusJSON(503, gin.H{
				"error": "too many request",
			})
		}
		
		ctx.Next()
	}
}
