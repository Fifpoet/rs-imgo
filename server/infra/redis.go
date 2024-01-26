package infra

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"rs-imgo/util"
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
		log.Fatal("ZAddBatchPng error")
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
		b64 := util.PathToB64(imgPaths[i])
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

func QueryPngByField(key, field string) string {
	res, err := rc.HGet(context.Background(), key, field).Result()
	if err == redis.Nil {
		log.Printf("QueryPngByField dict key not exisit")
	} else if err != nil {
		log.Fatal("QueryPngByField error", err)
	}
	return res
}

func HAddPng(key string, imgPath string, field string) int64 {
	b64 := util.PathToB64(imgPath)
	res, err := rc.HSet(context.Background(), key, field, b64).Result()
	if res > 0 {
		log.Printf("HAddPng 成功缓存{%v}张图片, key:{%v}, imgPath:{%v}", len(imgPath), key, imgPath)
	}
	if err != nil {
		log.Fatal("HAddPng错误")
	}
	return res
}
