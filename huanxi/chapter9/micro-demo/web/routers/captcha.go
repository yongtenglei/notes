package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/rey/micro-demo/dao/redis"
	"github.com/rey/micro-demo/log"
)

const (
	CaptchaExpiredTime = 120 * time.Second
)

func CaptchaHandler(c *gin.Context) {
	mobile, ok := c.GetQuery("mobile")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return

	}

	ds := captcha.RandomDigits(captcha.DefaultLen)
	img := captcha.NewImage("", ds, captcha.StdWidth, captcha.StdHeight)

	// chaptcha content
	cpc := ""
	for _, d := range ds {
		cpc += fmt.Sprintf("%d", d)
	}

	fmt.Println(cpc)

	redis.RedisClient.Set(context.Background(), mobile, cpc, CaptchaExpiredTime)

	var buf bytes.Buffer
	_, err := img.WriteTo(&buf)
	if err != nil {
		s := fmt.Sprintln("GenerateTokenChatcha WriteTo File failed: ", err)
		log.Logger.Info(s)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	img_base64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"msg":     "",
		"captcha": img_base64,
	})

}

// 写入文件再从文件得到编码(不提倡)
//func GenerateChatchaThroughFile() (err error) {
//fileName := "chaptcha.png"
//file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
//if err != nil {
//s := fmt.Sprintln("GenerateTokenChatcha OpenFile failed: ", err)
//log.Logger.Info(s)
//fmt.Println(err)
//return
//}
//defer file.Close()

//ds := captcha.RandomDigits(captcha.DefaultLen)
//img := captcha.NewImage("", ds, captcha.StdWidth, captcha.StdHeight)

//_, err = img.WriteTo(file)
//if err != nil {
//s := fmt.Sprintln("GenerateTokenChatcha WriteTo File failed: ", err)
//log.Logger.Info(s)
//fmt.Println(err)
//return
//}

//fmt.Println("===========", ds)

//img_base64, err := ToBase64(fileName)
//if err != nil {
//s := fmt.Sprintln("GenerateTokenChatcha ToBase64 failed: ", err)
//log.Logger.Info(s)
//fmt.Println(err)
//return
//}

//fmt.Println(img_base64)
//fmt.Println(err)
//return

//}

//func ToBase64(fileName string) (string, error) {
//fd, err := ioutil.ReadFile(fileName)
//if err != nil {
//return "", err
//}

//return base64.StdEncoding.EncodeToString(fd), nil
//}
