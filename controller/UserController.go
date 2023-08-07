package controller

import (
	"ginvueDome/common"
	"ginvueDome/model"
	"ginvueDome/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := common.InitDB()
	//获取参数
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	//数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须是11位",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, phone, password)
	//判断手机号是否存在
	if model.IsPhoneExist(db, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号已存在",
		})
		return
	}
	//创建用户
	user := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	db.Create(&user)

	//返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}
