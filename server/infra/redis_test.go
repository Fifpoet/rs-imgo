package infra

import (
	"fmt"
	"log"
	"os"
	"rs-imgo/global"
	"testing"
)

func setup() {
	fmt.Println("redis_test setup")
	InitRedis()
}

// 测试函数
func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	os.Exit(code)
}

func TestZAddBatchPng(t *testing.T) {
	key := global.ZsetKeyPrefix + "4"
	ZAddBatchPng(key, []string{"/home/fifpoet/Desktop/go_project/rs-imgo/server/static/output/0/1/2/0120.png"}, []int{1111})
}

func TestQueryPngByScore(t *testing.T) {
	key := global.ZsetKeyPrefix + "4"
	score := "1221"
	pngs := QueryPngByScore(key, score)
	log.Printf("pngs: {%v}", pngs)
}