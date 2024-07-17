/*
@File    :   router.go
@Time    :   2024/04/09 21:22:29
@Author  :   Luis
@Contact :   luis9527@163.com
*/

package controller

import (
	"github.com/gin-gonic/gin"
)

// 注册路由
func RegisterRouter(r *gin.Engine) {
	r.GET("/testapi", TestApi)
	r.POST("/api/login", Login.Auth)
	rgroup := r.Group("/api/k8s")
	rgroup.
	//工作流
	GET("/workflows", Workflow.GetList).
	GET("/workflow/detail", Workflow.GetById).
	POST("/workflow/create", Workflow.Create).
	DELETE("/workflow/del", Workflow.DelById).
	//pod操作
	GET("/pods", Pod.GetPods).
	GET("/pod/detail", Pod.GetPodDetail).
	DELETE("/pod/del", Pod.DeletePod).
	PUT("/pod/update", Pod.UpdatePod).
	GET("/pod/container", Pod.GetPodContainer).
	GET("/pod/log", Pod.GetPodLog).
	GET("/pod/numnp", Pod.GetPodNumPerNp).
	//deployment操作
	GET("/deployments", Deployment.GetDeployments).
	GET("/deployment/detail", Deployment.GetDeploymentDetail).
	PUT("/deployment/scale", Deployment.ScaleDeployment).
	DELETE("/deployment/del", Deployment.DeleteDeployment).
	PUT("/deployment/restart", Deployment.RestartDeployment).
	PUT("/deployment/update", Deployment.UpdateDeployment).
	GET("/deployment/numnp", Deployment.GetDeployNumPerNp).
	POST("/deployment/create", Deployment.CreateDeployment).
	//daemonset操作
	GET("/daemonsets", DaemonSet.GetDaemonSets).
	GET("/daemonset/detail", DaemonSet.GetDaemonSetDetail).
	DELETE("/daemonset/del", DaemonSet.DeleteDaemonSet).
	PUT("/daemonset/update", DaemonSet.UpdateDaemonSet).
	//statefulset操作
	GET("/statefulsets", StatefulSet.GetStatefulSets).
	GET("/statefulset/detail", StatefulSet.GetStatefulSetDetail).
	DELETE("/statefulset/del", StatefulSet.DeleteStatefulSet).
	PUT("/statefulset/update", StatefulSet.UpdateStatefulSet).
	//service操作
	GET("/services", Servicev1.GetServices).
	GET("/service/detail", Servicev1.GetServiceDetail).
	DELETE("/service/del", Servicev1.DeleteService).
	PUT("/service/update", Servicev1.UpdateService).
	POST("/service/create", Servicev1.CreateService).
	//ingress操作
	GET("/ingresses", Ingress.GetIngresses).
	GET("/ingress/detail", Ingress.GetIngressDetail).
	DELETE("/ingress/del", Ingress.DeleteIngress).
	PUT("/ingress/update", Ingress.UpdateIngress).
	POST("/ingress/create", Ingress.CreateIngress).
	//configmap操作
	GET("/configmaps", ConfigMap.GetConfigMaps).
	GET("/configmap/detail", ConfigMap.GetConfigMapDetail).
	DELETE("/configmap/del", ConfigMap.DeleteConfigMap).
	PUT("/configmap/update", ConfigMap.UpdateConfigMap).
	//sercret操作
	GET("/secrets", Secret.GetSecrets).
	GET("/secret/detail", Secret.GetSecretDetail).
	DELETE("/secret/del", Secret.DeleteSecret).
	PUT("/secret/update", Secret.UpdateSecret).
	//pvc操作
	GET("/pvcs", Pvc.GetPvcs).
	GET("/pvc/detail", Pvc.GetPvcDetail).
	DELETE("/pvc/del", Pvc.DeletePvc).
	PUT("/pvc/update", Pvc.UpdatePvc).
	//node操作
	GET("/nodes", Node.GetNodes).
	GET("/node/detail", Node.GetNodeDetail).
	//namespace操作
	GET("/namespaces", Namespace.GetNamespaces).
	GET("/namespace/detail", Namespace.GetNamespaceDetail).
	DELETE("/namespace/del", Namespace.DeleteNamespace).
	//pv操作
	GET("/pvs", Pv.GetPvs).
	GET("/pv/detail", Pv.GetPvDetail)
}
