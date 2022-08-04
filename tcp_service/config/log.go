package config

import (
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

var logger *log.Logger

// 必须实现Log函数,哪怕是个空函数
func Log(v ...interface{}) {
	logger.Println(v...)
}

// 初始化Log
func LogInit() {
	log_dir := LogDir
	// 创建日志文件
	file_name := time.Now().Format("2006-01-02")
	create_log_file(log_dir, file_name)

	// 定时每天零点创建新的日志文件
	create_file_timer := cron.New()
	EntryID, err := create_file_timer.AddFunc("0 0 * * *", func() {
		file_name = time.Now().Format("2006-01-02")
		create_log_file(log_dir, file_name)
	})
	if err != nil {
		println("ERROR: 创建日志文件失败", "\r\n\tEntryID: ", EntryID, "\t错误信息: ", err)
	}
	create_file_timer.Start()
}

func create_log_file(log_dir string, file_name string) {
	_, err := os.Stat(log_dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(log_dir, 0644)
		if err != nil {
			println("ERROR: 初始化日志文件失败\t", err)
		}
	}
	file := log_dir + file_name + ".log"
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		println("ERROR: 打开日志文件异常")
	} else {
		println("初始化日志文件成功")
	}
	logger = log.New(logFile, "success", log.Ldate|log.Ltime|log.Lshortfile)
}
