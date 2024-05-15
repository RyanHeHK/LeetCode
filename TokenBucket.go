package main

import (
	"time"
)

type TokenBucket struct {
	rate       int   // 令牌放入速率（单位：次/秒）
	capacity   int   // 令牌桶容量
	tokens     int   // 当前令牌数量
	lastUpdate int64 // 上次更新时间（毫秒）
}

func NewTokenBucket(rate, capacity int) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastUpdate: time.Now().UnixNano() / int64(time.Millisecond),
	}
}

func (tb *TokenBucket) Allow() bool {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := now - tb.lastUpdate
	tb.lastUpdate = now

	tb.tokens += int(elapsed) * tb.rate / 1000
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

//func main() {
//	tokenBucket := NewTokenBucket(10, 5) // 令牌放入速率为10次/秒，容量为5
//	for i := 0; i < 15; i++ {
//		if tokenBucket.Allow() {
//			fmt.Println("Request", i+1, "is allowed")
//		} else {
//			fmt.Println("Request", i+1, "is rejected")
//		}
//		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
//	}
//}
