/*
@File    :   init.go
@Time    :   2024/04/09 21:58:58
@Author  :   Luis
@Contact :   luis9527@163.com
*/

package service

import (
	"fmt"
	"k8s-server/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8sClientSet *kubernetes.Clientset

func InitK8sClientSet() {
	conf, err := clientcmd.BuildConfigFromFlags("", config.Config.GetString("Kubenertes.config"))
	if err != nil {
		fmt.Println("创建k8s配置失败, " + err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		fmt.Println("创建k8s clientSet失败, " + err.Error())
	} else {
		fmt.Println("创建k8s clientSet成功")

		K8sClientSet = clientSet
	}
}
