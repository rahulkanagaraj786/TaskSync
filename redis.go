package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Redis keys for storing jobs
var scheduledJobsKey = "scheduledJobs"
var backupJobsKey = "backupJobs"

// Redis client configuration
var redisClient = redis.NewClient(&redis.Options{
	Addr:       "localhost:6379", // Redis server address
	PoolSize:   5,                // Max connection pool size
	MaxRetries: 2,                // Maximum retry attempts
	DB:         0,                // Default Redis database
})

// Remove a job from the specified Redis set
func removeJobFromRedis(setName, jobName string) {
	err := redisClient.ZRem(setName, jobName).Err()
	if err != nil {
		fmt.Printf("[ERROR] Failed to remove job '%s' from Redis: %v\n", jobName, err)
	}
}

// Add a job to the specified Redis set with a given timestamp
func addJobToRedis(setName string, timestamp float64, jobName string) {
	err := redisClient.ZAdd(setName, redis.Z{
		Score:  timestamp, // Job timestamp as score
		Member: jobName,   // Job name as member
	}).Err()
	if err != nil {
		fmt.Printf("[ERROR] Failed to add job '%s' to Redis: %v\n", jobName, err)
	}
}

// Fetch jobs from the backup set with a timestamp less than or equal to the given timestamp
func findJobsByTime(timestamp float64) []redis.Z {
	jobs, err := redisClient.ZRangeByScoreWithScores(backupJobsKey, redis.ZRangeBy{
		Min: "0",
		Max: fmt.Sprintf("%f", timestamp),
	}).Result()
	if err != nil {
		fmt.Println("[ERROR] Failed to fetch jobs from Redis:", err)
	}
	return jobs
}

// Fetch the next scheduled job from the Redis set
func fetchNextJob() redis.ZWithKey {
	job, err := redisClient.BZPopMin(3*time.Second, scheduledJobsKey).Result()
	if err != nil {
		fmt.Printf("[ERROR] Failed to fetch the next job: %v\n", err)
	}
	return job
}
