/*
@File    :   testapi.go
@Time    :   2024/04/09 21:31:05
@Author  :   Luis
@Contact :   luis9527@163.com
*/

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestApi(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"msg":"test success!",
			"data":nil,
		})
}