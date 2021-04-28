package main

import (
	"github.com/gin-gonic/gin"
	"golang-admin/global"
	"golang-admin/internal/routers"
	"golang-admin/pkg/setting"
	"log"
	"net/http"
	"time"
)

//它的执行顺序是：全局变量初始化 => init 方法 => main 方法，但并不是建议滥用，
//因为如果 init 过多，你可能会迷失在各个库的 init 方法中，会非常麻烦。
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("setup setting err! err is %v", err)
	}

	err = setupDatabase()
	if err != nil {
		log.Fatalf("setup database err! err is %v", err)
	}

	err = setting.CreatTables(global.DBEngine)
	if err != nil {
		log.Fatalf("create tables err! err is %v", err)
	}
}

// @title yd的golang学习后台
// @version 1.0
// @description Go Go Go！！！
// @termsOfService https://github.com/yd2008/golang-admin
func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Third", &global.ThirdSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("ALiOSS", &global.ALiOSSSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDatabase() error {
	DBEngine, err := setting.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DBEngine = DBEngine
	return nil
}
