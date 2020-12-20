package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description 练习Go
// @termsOfService https:google.com
func main() {
	router := routers.NewRouter()
	s := http.Server{
		Addr:           ":" + global.ServeSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServeSetting.ReadTimeout,
		WriteTimeout:   global.ServeSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Infof("%s: go-programming-tour-book/%s", "qyu", "blog-service")

	_ = s.ListenAndServe()

}

func setupSetting() error {
	setg, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setg.ReadSection("Server", &global.ServeSetting)
	if err != nil {
		return err
	}
	err = setg.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setg.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServeSetting.ReadTimeout *= time.Second
	global.ServeSetting.WriteTimeout *= time.Second

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +
			global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupValidator() error {
	// 自定义validator设置
	binding.Validator = global.Validator
	return nil
}
