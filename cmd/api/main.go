package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"teacupapi/config"
	"teacupapi/db"
	"teacupapi/libs"
	"teacupapi/libs/ginValidator"
	glog "teacupapi/logs"
	"teacupapi/models"
	apiRouter "teacupapi/router"
	"time"
)

var (
	confPath = flag.String("config", "../../config/app.dev.ini", "profilePath")
	httpSrv  *http.Server
)

// Init main init
func Init() error {
	flag.Parse()

	err := config.InitConfig(confPath)
	if err != nil {
		return fmt.Errorf("init config is err: %v", err)
	}

	err = glog.InitLog()
	if err != nil {
		return fmt.Errorf("init log is err: %v", err)
	}

	//err = natsUtils.Init()
	//if err != nil {
	//	return fmt.Errorf("init nats streaming is err: %v", err)
	//}

	//err = apns.Init()
	//if err != nil {
	//	return fmt.Errorf("init apns  is err: %v", err)
	//}
	return nil
}

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			tmpStr := fmt.Sprintf("err: %v, panic==> %s\n", err, string(buf[:n]))
			//glog.Errorf(tmpStr)
			fmt.Println(tmpStr)
		}
	}()

	err := Init()
	if err != nil {
		fmt.Println("main init err: ", err)
		return
	}

	err = db.InitDB()
	if err != nil {
		glog.Error("init db err: ", err)
		return
	}
	defer db.Close()

	err = libs.Initlibs()
	if err != nil {
		glog.Errorf("init libs err: %v", err)
		return
	}

	//gin
	engine := gin.New()
	engine.StaticFS("/images", http.Dir(config.GetUploadConf().FileStoragePath))
	corsConfig := cors.DefaultConfig() // 解决跨域问题
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders(models.ClientType, models.HeaderToken, models.Version)
	engine.Use(cors.New(corsConfig), gin.Recovery())
	validatorRegister()
	apiRouter.ApiRouter(engine)

	baseConf := config.GetBaseConf()
	// http services
	httpSrv = &http.Server{
		Addr:    ":" + baseConf.HttpPort,
		Handler: engine,
	}

	if baseConf.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go handleSignal(quit)

	glog.Infof("dev env=%v is start", baseConf.Env)
	// http service connections
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		glog.Errorf("http listen: %v", err)
		panic(err)
	}
}

func handleSignal(c chan os.Signal) {
	sigValue := <-c
	switch sigValue {
	case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
		fmt.Printf("Shutdown quickly sig=%v, bye...\n", sigValue)
		glog.Infof("Shutdown quickly sig=%v, bye...", sigValue)
	case syscall.SIGHUP:
		fmt.Printf("Shutdown gracefully sig=%v, bye...\n", sigValue)
		glog.Infof("Shutdown gracefully sig=%v, bye...", sigValue)
		// do graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpSrv.Shutdown(ctx); err != nil {
			glog.Error("http Server Shutdown err:", err)
		}
	}
	os.Exit(0)
}

func validatorRegister() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("trimStr", ginValidator.TrimHeaderTailBlank)
		if err != nil {
			glog.Errorf("注册validator校验器，err:%+v", err)
		}
	}
}
