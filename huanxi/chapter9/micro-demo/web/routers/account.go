package routers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/rey/micro-demo/log"
	"github.com/rey/micro-demo/proto/account"
	"github.com/rey/micro-demo/web/res"
)

func AccountListHandler(c *gin.Context) {
	// 获取分页信息
	page_str := c.DefaultQuery("page", "1")
	pageSize_str := c.DefaultQuery("pagesize", "5")
	page, _ := strconv.Atoi(page_str)
	pageSize, _ := strconv.Atoi(pageSize_str)

	// 拨号
	conn, err := grpc.Dial("localhost:9409", grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler dial filed: %+v", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 创建Client, 调用grpc函数
	ac := account.NewAccountServiceClient(conn)
	listRes, err := ac.GetAccountList(context.Background(), &account.PagingReq{
		PageNo:   uint32(page),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		s := fmt.Sprintf("AccountListHandler GetAccountList filed: %+v", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 返回
	var accountList []res.AccountRes
	for _, account := range listRes.AccountList {
		accountList = append(accountList, *res.AccountInfo2AccountRes(account))
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "",
		"total": listRes.Total,
		"data":  accountList,
	})
	return
}
