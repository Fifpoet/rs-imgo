package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"path/filepath"
	"rs-imgo/global"
	"rs-imgo/util"
	"strings"
)

var db *sql.DB

func InitMysql() {
	msdb, err := sql.Open("mysql", "root:zc1669600787@tcp(localhost:3306)/footprint")
	if err != nil {
		log.Fatal("mysql conn failed", err.Error())
	}
	log.Printf("启动mysql成功")
	db = msdb
}

func UploadImg() {
	base := global.ImgPath
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// process xxx.png
			key := strings.Split(info.Name(), ".")[0]
			b64 := util.PathToB64(path)
			insert(key, b64)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录 %s 失败: %v\n", base, err)
		return
	}
}

func insert(key, b64 string) {
	stmt, _ := db.Prepare("INSERT INTO tiles(tile_id, pic) VALUES(?, ?)")
	defer stmt.Close()
	_, err := stmt.Exec(key, b64)
	if err != nil {
		fmt.Println("插入记录失败:", err.Error())
		return
	}
}

func QueryB64ByQuadKey(key string) string {
	var res string
	err := db.QueryRow("SELECT pic FROM tiles WHERE tile_id = ?", key).Scan(&res)
	if err != nil {
		log.Printf("mysql query failed")
	}
	return res
}
