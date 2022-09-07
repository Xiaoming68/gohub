package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gohub/initPack/logger"
	"sync"
	"time"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var Redis *RedisClient

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, username string, password string, db int) *RedisClient {

	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	// 使用默认的 context
	rds.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}
