package service

import (
	"github.com/gin-gonic/gin"
	"im/helper"
	models "im/models"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "账号或密码不能为空",
		})
		return
	}
	ub, err := models.GetUserBasicByAccountPassword(account, helper.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "账号或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "系统错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(c *gin.Context) {
	user, _ := c.Get("user_claims")
	userClaims := user.(*helper.UserClaims)
	userBasic, err := models.GetUserBasicByIdentity(userClaims.Identity)
	if err != nil {
		log.Printf("[DB ERROR]:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": userBasic,
	})
}
