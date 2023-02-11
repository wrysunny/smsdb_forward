package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var (
	db *sql.DB
)

func init() {
	// log日志设置
	logFile, err := os.OpenFile(fmt.Sprintf("/home/javashell/sms/logs/run-%s.log", time.Now().Format("20060102")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Fatalln("open/create log file failed:", err.Error())
	}
	log.SetPrefix("[Info] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
	// 打开数据库
	db, err = sql.Open("sqlite3", "/home/javashell/sms/sms.db")
	if err != nil {
		log.SetPrefix("[Error] ")
		log.Fatalln("open sqlite3 database failed:", err.Error())
	}
	// 健康度检查
	health()
	// 查询未读消息
	query()

	// 关闭数据库连接
	db.Close()
}
