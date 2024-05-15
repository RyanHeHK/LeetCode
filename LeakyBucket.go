package main

import (
	"time"
)

type LeakyBucket struct {
	rate       int   // 漏桶流出速率（单位：次/秒）
	capacity   int   // 漏桶容量
	water      int   // 漏桶当前水量
	lastLeakMs int64 // 上次漏水时间（毫秒）
}

func NewLeakyBucket(rate, capacity int) *LeakyBucket {
	return &LeakyBucket{
		rate:       rate,
		capacity:   capacity,
		water:      0,
		lastLeakMs: time.Now().UnixNano() / int64(time.Millisecond),
	}
}

func (lb *LeakyBucket) Allow() bool {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := now - lb.lastLeakMs
	lb.lastLeakMs = now

	leakAmount := int(elapsed) * lb.rate / 1000
	lb.water -= leakAmount
	if lb.water < 0 {
		lb.water = 0
	}

	if lb.water < lb.capacity {
		lb.water++
		return true
	}

	return false
}

//func main() {
//	leakyBucket := NewLeakyBucket(10, 5) // 漏桶流出速率为10次/秒，容量为5
//	for i := 0; i < 15; i++ {
//		if leakyBucket.Allow() {
//			fmt.Println("Request", i+1, "is allowed")
//		} else {
//			fmt.Println("Request", i+1, "is rejected")
//		}
//		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
//	}
//}
