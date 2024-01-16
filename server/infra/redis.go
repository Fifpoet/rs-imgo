package infra

import (
	"context"
	"encoding/base64"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var rc *redis.Client

func InitRedis() {
	c := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := c.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("启动redis失败", err.Error())
	}
	log.Printf("启动redis成功")
	rc = c
}

func ZAddBatchPng(key string, imgPaths []string, scores []int) int64 {
	zs := genZAddReq(imgPaths, scores)
	res, err := rc.ZAddNX(context.Background(), key, zs...).Result()
	if res > 0 {
		log.Printf("ZAddBatchPng成功缓存{%v}张图片, key:{%v}, imgPaths:{%v}, scores:{%v}", len(imgPaths), key, imgPaths, scores)
	}
	if err != nil {
		log.Fatal("xxx")
	}
	return res
}

func genZAddReq(imgPaths []string, scores []int) []redis.Z {
	l := len(imgPaths)
	if l != len(scores) {
		panic("param error")
	}
	res := make([]redis.Z, l)

	for i := 0; i < l; i++ {
		bs, err := os.ReadFile(imgPaths[i])
		if err != nil {
			panic(err)
		}
		b64 := base64.StdEncoding.EncodeToString(bs)
		res[i] = redis.Z{Score: float64(scores[i]), Member: b64}
	}
	return res

}

func QueryPngByScore(key string, score string) []string {
	pngs, err := rc.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{
		Min: score,
		Max: score,
	}).Result()
	if err != nil {
		log.Fatal("QueryPngByScore错误")
	}
	log.Printf("QueryPngByScore: req {key: %v, score: %v} resLen {%v}", key, score, len(pngs))
	return pngs
}
