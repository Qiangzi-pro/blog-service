package main

import (
	"context"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/go-programming-tour-book/blog-service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	var (
		err error
	)
	err = setupSetting()
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
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

	global.Logger.Infof(context.Background(), "%s: go-programming-tour-book/%s", "qyu", "blog-service")

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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
