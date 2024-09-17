package redisdb

import (
	"context"
	"errors"
	"fmt"
	"nuryanto2121/cukur_in_web/pkg/setting"
	"strconv"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redislib "github.com/redis/go-redis/v9"
)

// var rdb *redislib.Client

type RedisHandler struct {
	client *redislib.Client
	pool   *redis.Pool
	rs     *redsync.Redsync
}

// Setup :
// func Setup() {
// 	now := time.Now()
// 	conString := fmt.Sprintf("%s:%d", setting.FileConfigSetting.RedisDBSetting.Host, setting.FileConfigSetting.RedisDBSetting.Port)
// 	rdb = redislib.NewClient(&redislib.Options{
// 		Addr:     conString,
// 		Password: setting.FileConfigSetting.RedisDBSetting.Password,
// 		DB:       setting.FileConfigSetting.RedisDBSetting.DB,
// 	})
// 	_, err := rdb.Ping(context.Background()).Result()
// 	if err != nil {
// 		fmt.Println(err)
// 		// logging.Error("0", err)
// 		// logging.Fatal("0", err)
// 	}
// 	// fmt.Println("Mem Cache is Ready...")

// 	timeSpent := time.Since(now)
// 	log.Printf("Config redis is ready in %v", timeSpent)
// }

var ErrMutexAlreadyExist = errors.New("mutex with same key already exist")

func New() *RedisHandler {
	conString := fmt.Sprintf("%s:%d", setting.FileConfigSetting.RedisDBSetting.Host, setting.FileConfigSetting.RedisDBSetting.Port)
	clientOptions := &redislib.Options{
		Addr: conString,
		DB:   setting.FileConfigSetting.RedisDBSetting.DB,
	}
	if setting.FileConfigSetting.RedisDBSetting.Password != "" {
		clientOptions.Password = setting.FileConfigSetting.RedisDBSetting.Password
	}
	client := redislib.NewClient(clientOptions)
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	return &RedisHandler{client: client, rs: rs}
}

func (r *RedisHandler) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
func (r *RedisHandler) NewMutex(ctx context.Context, key string, timeout time.Duration) (*redsync.Mutex, error) {
	exists, err := r.Exists(ctx, key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrMutexAlreadyExist
	}
	return r.rs.NewMutex(key, redsync.WithExpiry(timeout)), nil
}
func (r *RedisHandler) CreateAndLockMutex(ctx context.Context, key string, timeout time.Duration) (*redsync.Mutex, error) {
	mutex, err := r.NewMutex(ctx, key, timeout)
	if err != nil {
		return nil, err
	}
	if err = mutex.TryLockContext(ctx); err != nil {
		return nil, err
	}
	return mutex, nil
}
func (r *RedisHandler) GetInt(ctx context.Context, key string) (int, error) {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(redislib.Nil, err) {
		return 0, nil
	}
	return strconv.Atoi(val)
}
func (r *RedisHandler) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
func (r *RedisHandler) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
func (r *RedisHandler) IncrWithTTL(ctx context.Context, key string, ttl time.Duration) error {
	if err := r.client.Incr(ctx, key).Err(); err != nil {
		return err
	}
	r.client.Expire(ctx, key, ttl)
	return nil
}
