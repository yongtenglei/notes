package routers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rey/micro-demo/log"
	myjwt "github.com/rey/micro-demo/pkg/jwt"
	"github.com/rey/micro-demo/proto/account"
	"github.com/rey/micro-demo/web/req"
	"google.golang.org/grpc"
)

func LoginByPasswordHandler(c *gin.Context) {
	// 获取登录信息
	var loginByPasswordParam req.LoginByPasswordParam
	err := c.ShouldBindJSON(&loginByPasswordParam)
	if err != nil {
		s := fmt.Sprintln("LoginByPasswordHandler ShouldBindJSON failed", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 进行密码校验
	conn, err := grpc.Dial("localhost:9409", grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler dial filed: %+v", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	sc := account.NewAccountServiceClient(conn)

	mobileRes, err := sc.GetAccountByMobile(context.Background(), &account.MobileReq{
		Mobile: loginByPasswordParam.Mobile,
	})
	if err != nil {
		s := fmt.Sprintf("AccountListHandler GetAccountByMobile filed: %+v", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	checkRes, err := sc.CheckAccountByID(context.Background(), &account.CheckAccountByIDReq{
		Id:       mobileRes.Account.Id,
		Password: loginByPasswordParam.Password,
	})
	if err != nil {
		s := fmt.Sprintf("AccountListHandler CheckAccountByID filed: %+v", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	loginRes := "登录成功"
	// 没有验证通过
	if !checkRes.Ok {
		loginRes = "登录失败"
		c.JSON(http.StatusOK, gin.H{
			"msg": loginRes,
		})
		return
	}

	// 验证成功
	j := myjwt.NewJWT()
	claims := myjwt.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		},
		ID:       uint64(mobileRes.Account.Id),
		UserName: mobileRes.Account.Username,
	}
	token, err := j.GenerateToken(claims)
	if err != nil {
		s := fmt.Sprintf("AccountListHandler GenerateToken filed: %+v", err.Error())
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   loginRes,
		"token": token,
	})
}
