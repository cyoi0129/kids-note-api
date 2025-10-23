package models

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	// "context"
	// "encoding/json"
	// "os"
	// "github.com/joho/godotenv"
	// "github.com/redis/go-redis/v9"
)

/**
	インメモリキャッシュを使って実装するコード
**/

var go_cache *cache.Cache

func CreateCache() {
	c := cache.New(720*time.Hour, 1440*time.Hour)
	go_cache = c
}

func SetCache(k string, v interface{}) { // キャッシュを保存
	go_cache.Set(k, v, 720*time.Hour) // 保存期間を指定
}

func GetCache(k string) (interface{}, error) { // キャッシュを取得
	err := errors.New("no cache")
	r, ok := go_cache.Get(k)
	if ok {
		err = nil
	}
	return r, err
}

func DeleteCache(k string) { // キャッシュを削除
	go_cache.Delete(k)
}

/**
	キャッシュはRedisを使って実装するコード
**/

// var go_cache *redis.Client
// var ctx = context.Background()

// func CreateCache() {
// 	godotenv.Load(".env")
// 	godotenv.Load()

// 	go_cache = redis.NewClient(&redis.Options{
// 		Addr:     os.Getenv("CACHE_URL"), // redisの接続先
// 		Password: "",                     // パスワードなし
// 		DB:       0,                      // デフォルトDBを使う
// 		PoolSize: 1000,                   // サイズの上限
// 	})
// 	if err := go_cache.Ping(ctx).Err(); err != nil {
// 		panic(err)
// 	}
// 	go_cache.FlushAll(ctx)
// }

// func SetCache(k string, v interface{}) (interface{}, error) { // キャッシュを保存
// 	data, err := json.Marshal(v) // オブジェクト式をラッパーしたjson形式へ変換
// 	result, err := go_cache.Set(ctx, k, data, 0).Result()
// 	go_cache.Expire(ctx, k, 3*time.Minute) // 保存期間を指定
// 	return result, err
// }

// func GetCache(k string) (interface{}, error) { // キャッシュを取得
// 	var data interface{}
// 	result, err := go_cache.Get(ctx, k).Result()
// 	json.Unmarshal([]byte(result), &data) // キャッシュに保存されたjson形式へまたオブジェクトへ変換
// 	return data, err
// }

// func DeleteCache(k string) (interface{}, error) { // キャッシュを削除
// 	result, err := go_cache.Del(ctx, k).Result()
// 	return result, err
// }
