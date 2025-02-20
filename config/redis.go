package config

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal("Invalid Redis URL:", err)
	}

	// ✅ Skip certificate verification (Fix for Fly.io + Upstash)
	opt.TLSConfig = &tls.Config{
		InsecureSkipVerify: true, // 🔥 This line disables cert verification
	}

	RedisClient = redis.NewClient(opt)

	// Test Redis connection
	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis (Upstash)")
}
