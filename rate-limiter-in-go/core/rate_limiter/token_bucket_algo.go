package rate_limiter

import (
	"errors"
	"time"
)

type Bucket struct {
	Capacity           int // Bucket max capacity
	Tokens             int
	RefillRateTokens   int   // Refill `x` tokens
	RefillRateDuration int64 // Refill `x` tokens per `y` duration
	LastRefillTime     int64
}

func NewBucket(capacity, tokens, refillRateTokens int) *Bucket {
	return &Bucket{
		Capacity:           capacity,
		Tokens:             tokens,
		RefillRateTokens:   refillRateTokens,
		RefillRateDuration: 1000 * time.Millisecond.Milliseconds(), // 1 Second
		LastRefillTime:     time.Now().UnixMilli(),
	}
}

func (b *Bucket) TakeoutToken() error {
	currentTime := time.Now().UnixMilli()
	diff := currentTime - b.LastRefillTime

	if diff >= b.RefillRateDuration {
		// If the diff between current time and last refill time exceeds refile rate duration
		b.RefillBucket()
		b.LastRefillTime = currentTime
	}

	if b.Tokens <= 0 {
		return errors.New("no tokens. limit exceeded")
	}

	b.Tokens -= 1
	return nil
}

func (b *Bucket) RefillBucket() {
	newTokens := b.Tokens + b.RefillRateTokens
	if newTokens > b.Capacity {
		b.Tokens += newTokens - b.Capacity // Assigning to `b.Capacity` should also be fine
	} else {
		b.Tokens = newTokens
	}
}

func TokenBucketInit() map[string]*Bucket {
	return make(map[string]*Bucket)
}

func TokenBucketAlgo(buckets *map[string]*Bucket) func(string) error {
	return func(userIdentifier string) error {
		bucket, ok := (*buckets)[userIdentifier]

		if !ok {
			// For testing
			bucket = NewBucket(1, 1, 1)
			(*buckets)[userIdentifier] = bucket
		}

		err := bucket.TakeoutToken()
		if err != nil {
			return err
		}

		return nil
	}
}
