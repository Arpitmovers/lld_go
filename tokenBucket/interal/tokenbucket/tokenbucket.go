package tokenbucket

// https://medium.com/@oneconfusedindian/decoding-complexity-token-bucket-algorithm-ba80885acacd

// https://dev.to/leapcell/build-a-token-bucket-limiter-in-go-in-under-100-lines-4f61

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity int
	rate     int //  tokens to add per second
	tokens   int // current tokens
	mutex    sync.Mutex
	lastFill time.Time
}

func NewTokenBucket(capacity int, rate int) *TokenBucket {
	tb := &TokenBucket{
		capacity: capacity, // oversall capacity
		rate:     rate,
		tokens:   capacity,   // current tokens that we hold
		lastFill: time.Now(), // time when lastFill happend
	}
	go tb.refillTokens()
	return tb
}

// function to refill TOkens, based on
func (tb *TokenBucket) refillTokens() {

	currTime := time.Now()
	timeDiff := currTime.Sub(tb.lastFill) // time difference

	tokenToAdd := int(timeDiff * time.Duration(tb.rate))

	fmt.Println("tokenToAdd", tokenToAdd)

	if tokenToAdd > 0 {
		tb.tokens = min(tokenToAdd+tb.tokens, tb.capacity)
	}

}

// attempts to take x no of tokens from bucket, return bool based on task
func (tb *TokenBucket) acquireToken(count int) bool {
	tb.mutex.Lock()

	defer tb.mutex.Unlock()
	tb.refillTokens()
	if count < tb.tokens {
		tb.tokens = tb.tokens - count
		return true
	}

	return false
}
