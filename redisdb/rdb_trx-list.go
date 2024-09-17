package redisdb

import (
	"context"
	"fmt"
	"time"
)

// GetList :
func (r *RedisHandler) GetList(ctx context.Context, key string) ([]string, error) {
	list, err := r.client.SMembers(ctx, key).Result()
	return list, err
}

// RemoveList :
func (r *RedisHandler) RemoveList(ctx context.Context, key string, val interface{}) error {
	_, err := r.client.SRem(ctx, key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// AddList :
func (r *RedisHandler) AddList(ctx context.Context, key, val string) error {
	_, err := r.client.SAdd(ctx, key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// TurncateList :
func (r *RedisHandler) TurncateList(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

// AddSession :
func (r *RedisHandler) AddSession(ctx context.Context, key string, val interface{}, mn time.Duration) error {
	set, err := r.client.Set(ctx, key, val, mn).Result()
	if err != nil {
		return err
	}
	fmt.Println(set)
	return nil
}

// GetSession :
func (r *RedisHandler) GetSession(ctx context.Context, key string) interface{} {
	value := r.client.Get(ctx, key).Val()
	fmt.Println(value)
	return value
}
