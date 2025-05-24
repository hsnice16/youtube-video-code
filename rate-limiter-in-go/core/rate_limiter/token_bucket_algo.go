package rate_limiter

import (
	"fmt"
	"time"
)

type Bucket struct {
	Capacity int // Bucket max capacity
	Tokens     int
	RefillRateTokens int // Refill `x` tokens
	RefillRateDuration time.Duration // Refill `x` tokens per `y` duration
	LastRefillTime time.Time
}

func NewBucket(capacity, tokens, refillRateTokens int) Bucket {
	return Bucket{
		Capacity: capacity,
		Tokens:     tokens,
		RefillRateTokens: refillRateTokens,
		RefillRateDuration: 1 * time.Second,
		LastRefillTime: time.Now(),
	}
}

func (b *Bucket) TakeoutToken() {
	currentTime := time.Now()
	if(time.Duration(currentTime.Sub(b.LastRefillTime).Seconds()) >= b.RefillRateDuration) {
		
	}
}

func (b *Bucket) RefillBucket() {
	newTokens := b.Tokens + b.RefillRateTokens
	if newTokens > b.Capacity {
		b.Tokens += newTokens - b.Capacity
	} else {
		b.Tokens = newTokens
	}
}

func TokenBucketInit() map[string]Bucket {
	return make(map[string]Bucket)
}

func TokenBucketAlgo(buckets map[string]Bucket) func(string) error {
	return func(userIdentifier string) error {
		if bucket, ok := buckets[userIdentifier]; ok {

		} else {
			bucket := Bucket{Capacity:50, Tokens: 50, RefillRate: 5}
			buckets[userIdentifier] = bucket
		}

		fmt.Println("Inside Token Bucket Algo")
		return nil
	}
}
