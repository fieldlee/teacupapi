package redisdb

import (
	"github.com/go-redis/redis"
	"teacupapi/config"
	"time"
)

var RedisNil = redis.Nil

// RedisDb redis db
type RedisDb struct {
	wPool *redis.Client
	rPool *redis.Client
}

var redisDb *RedisDb

// InitRedis init redis
func InitRedis() error {
	wConf := config.GetRedisConf()
	wPool, err := initRedis(wConf.RedisHost, wConf.RedisAuth, wConf.RedisPoolSize)
	if err != nil {
		return err
	}

	rConf := config.GetRRedisConf()
	rPool, err := initRedis(rConf.RRedisHost, rConf.RRedisAuth, rConf.RRedisPoolSize)
	if err != nil {
		return err
	}

	redisD := &RedisDb{
		wPool: wPool,
		rPool: rPool,
	}
	redisDb = redisD

	return nil
}

// initRedisWrite  init redis write
func initRedis(host, auth string, poolSize int) (*redis.Client, error) {
	redisPool := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     auth,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     poolSize,
		PoolTimeout:  30 * time.Second,
		MaxRetries:   2,
		IdleTimeout:  5 * time.Minute,
	})

	_, err := redisPool.Ping().Result()
	if err != nil {
		return nil, err
	}

	return redisPool, nil
}

func getPool(isSlave bool) *redis.Client {
	if isSlave {
		return redisDb.rPool
	}

	return redisDb.wPool
}

// GetKey 根据key获取redis键值
func GetKey(isSlave bool, key string) (string, error) {
	value, err := getPool(isSlave).Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func GetTTL(isSlave bool, key string) (time.Duration, error) {
	value, err := getPool(isSlave).TTL(key).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

// GetKey 根据key获取redis键值, 返回 byte
func GetKeyBytes(isSlave bool, key string) ([]byte, error) {
	return getPool(isSlave).Get(key).Bytes()
}

// SetNotExpireKV 设置不过期的 key
func SetNotExpireKV(key, value string) error {
	err := redisDb.wPool.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetExpireKV 设置过期的 key
func SetExpireKV(key, value string, expire time.Duration) error {
	err := redisDb.wPool.Set(key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetExpireKey 设置 key 过期
func SetExpireKey(key string, expire time.Duration) error {
	err := redisDb.wPool.Expire(key, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetNX 设置 key, value 以及过期时间
func SetNX(key string, value string, expire time.Duration) (bool, error) {
	flag, err := redisDb.wPool.SetNX(key, value, expire).Result()
	if err != nil {
		return false, err
	}
	return flag, nil
}

// DelKey 删除 redis 的key
func DelKey(key string) error {
	return redisDb.wPool.Del(key).Err()
}

// KeyExist 判断某一个key 是否存在
func KeyExist(isSlave bool, keys string) (bool, error) {
	count, err := getPool(isSlave).Exists(keys).Result()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// HSet 设置 hash
func HSet(key, field string, value interface{}) error {
	return redisDb.wPool.HSet(key, field, value).Err()
}

// HMSet 批量存储 hash
func HMSet(key string, fields map[string]interface{}) error {
	if len(fields) < 1 {
		return nil
	}

	err := redisDb.wPool.HMSet(key, fields).Err()
	if err != nil {
		return err
	}

	return nil
}

// HGet 获取单个 hash
func HGet(isSlave bool, key, field string) (string, error) {
	return getPool(isSlave).HGet(key, field).Result()
}

// HMGet 批量获取 hash
func HMGet(isSlave bool, key string, fields ...string) ([]interface{}, error) {
	res, err := getPool(isSlave).HMGet(key, fields...).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// HScan 获取 hash 键值树
func HScan(isSlave bool, key string, curson uint64, match string, count int64) ([]string, uint64, error) {
	return getPool(isSlave).HScan(key, curson, match, count).Result()
}

// // HGetAll 获取整个 hash, 公司不允许用 hgetall 这个命令
// func HGetAll(key string) (map[string]string, error) {
// 	res, err := redisDb.redisPool.HGetAll(key).Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
func HLen(isSlave bool, key string) (int, error) {
	res, err := getPool(isSlave).HLen(key).Result()
	if err != nil {
		return 0, err
	}

	return int(res), nil
}

// HDel 删除 hash key
func HDel(key string, fields ...string) error {
	err := redisDb.wPool.HDel(key, fields...).Err()
	if err != nil {
		return err
	}

	return nil
}

// RPush 在名称为key的list尾添加一个值为value的元素
func RPush(key string, values ...interface{}) error {
	return redisDb.wPool.RPush(key, values...).Err()
}

// LPush 在名称为key的list头添加一个值为value的 元素
func LPush(key string, values ...interface{}) error {
	return redisDb.wPool.LPush(key, values...).Err()
}

// LLen 返回名称为key的list的长度
func LLen(key string, isSlave bool) (int64, error) {
	return getPool(isSlave).LLen(key).Result()
}

// LRange 返回名称为key的list中start至end之间的元素, start为0, end为-1 则是获取所有 list key
func LRange(isSlave bool, key string, start, end int64) ([]string, error) {
	return getPool(isSlave).LRange(key, start, end).Result()
}

// LSet 给名称为key的list中index位置的元素赋值
func LSet(key string, index int64, value interface{}) error {
	return redisDb.wPool.LSet(key, index, value).Err()
}

// LRem 删除count个key的list中值为value的元素
func LRem(key string, count int64, value interface{}) error {
	return redisDb.wPool.LRem(key, count, value).Err()
}

// ZAdd zset 有序集合中增加一个成员
func ZAdd(key, member string, score float64) error {
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	_, err := redisDb.wPool.ZAdd(key, z).Result()
	if err != nil {
		return err
	}

	return nil
}

// ZCount zset 有序集合中 min-max中的成员数量
func ZCount(isSlave bool, key, min, max string) (int64, error) {
	count, err := getPool(isSlave).ZCount(key, min, max).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ZCARD 获取 zset 中元素的数量
func ZCARD(isSlave bool, key string) (int64, error) {
	count, err := getPool(isSlave).ZCard(key).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ZRange 通过索引区间返回有序集合成指定区间内的成员
func ZRange(isSlave bool, key string, start, stop int64) ([]string, error) {
	arr, err := getPool(isSlave).ZRange(key, start, stop).Result()
	if err != nil {
		return []string{}, err
	}

	return arr, nil
}

// ZRangeByScore 通过索引区间返回有序集合成指定区间内的成员
func ZRangeByScore(isSlave bool, key string, min, max string) ([]string, error) {
	opt := redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	arr, err := getPool(isSlave).ZRangeByScore(key, opt).Result()
	if err != nil {
		return []string{}, err
	}

	return arr, nil
}

func ZRem(key string, members ...string) error {
	return redisDb.wPool.ZRem(key, members).Err()
}

func HGetBytesByField(isSlave bool, key, filed string) ([]byte, error) {
	return getPool(isSlave).HGet(key, filed).Bytes()
}

func SIsMember(isSlave bool, key, field string) (bool, error) {
	return getPool(isSlave).SIsMember(key, field).Result()
}
func Incr(key string) (int64, error) {
	return redisDb.wPool.Incr(key).Result()
}

func SMembers(isSlave bool, key string) ([]string, error) {
	return getPool(isSlave).SMembers(key).Result()
}

func SAdd(key string, members ...interface{}) (int64, error) {
	return redisDb.wPool.SAdd(key, members...).Result()
}

func SRem(key string, members ...interface{}) (int64, error) {
	return redisDb.wPool.SRem(key, members...).Result()
}

func LIndex(key string, index int64) (string, error) {
	return redisDb.wPool.LIndex(key, index).Result()
}
