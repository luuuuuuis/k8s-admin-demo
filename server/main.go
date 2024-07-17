/*
@File    :   main.go
@Time    :   2024/04/09 16:24:41
@Author  :   Luis
@Contact :   luis9527@163.com
*/

package main

import (
	"k8s-server/config"
	"k8s-server/controller"
	"k8s-server/db"
	"k8s-server/middleware"
	"k8s-server/service"
	"k8s-server/utils"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	utils.LogInit()
	// 初始化K8s clientset
	service.InitK8sClientSet()
	// 初始化数据库
	db.Init()
	// 创建gin实例
	r := gin.New()
	// 使用日志中间件
	r.Use(middleware.GinLogger, middleware.Cors())
	// 初始化路由
	controller.RegisterRouter(r)
	//终端websocket
	go func() {
		http.HandleFunc("/ws", service.Terminal.WsHandler)
		http.ListenAndServe(":8082", nil)
	}()
	// 运行程序
	err := r.Run(config.Config.GetString("Server.listenAddr"))
	if err != nil {
		utils.Logger.Error().
			Err(errors.New("启动失败")).
			Stack().
			Msg(err.Error())
	}
	//关闭db连接
	db.Close()
}
