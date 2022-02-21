package controller

import (
	"douban/dao"
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginByPwd(c *gin.Context) {
	phone := c.PostForm("phone")
	pwd := c.PostForm("password")

	err := service.LoginByPwd(phone, pwd)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	//获取token
	token, err := tool.CreateToken(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"token":  "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"token":  token,
	})
}

func Register(c *gin.Context) {
	var u model.User
	err := c.ShouldBind(&u)
	log.Println(u)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	err = service.Register(u)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"token":  "",
		})
		return
	}

	var tokenString string
	tokenString, err = tool.CreateToken(u.Phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"token":  "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"token":  tokenString,
	})
}

//CreateIntroduction 创建用户介绍
func CreateIntroduction(c *gin.Context) {
	phone, exists := c.Get("phone")
	if !exists {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	userIntroduction := c.PostForm("user_introduction")
	err := service.CreateIntroduction(phone.(string), userIntroduction)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

//CreateSign 创建用户签名
func CreateSign(c *gin.Context) {
	phone, exists := c.Get("phone")
	if !exists {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	sign := c.PostForm("user_sign")
	err := service.CreateSign(phone.(string), sign)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

//UploadUserAvatar 上传用户头像
func UploadUserAvatar(c *gin.Context) {
	//从上下文中获取用户标识
	phone, exists := c.Get("phone")
	if !exists {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	//获取头像文件
	fileHeader, err := c.FormFile("user_avatar")
	if err != nil {
		c.JSON(500, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	//拼接相对路径
	dst := fmt.Sprintf(".././douban/static/picture/useravatar/%s", fileHeader.Filename) //todo 抽离至配置文件
	//保存文件
	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	//保存图像名
	err = service.UploadUserAvatar(phone.(string), fileHeader.Filename)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

//GetInfo 获取用户部分信息
func GetInfo(c *gin.Context) {
	//从上下文中获取用户标识
	phone, exists := c.Get("phone")
	if !exists {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
			"user_info": "",
		})
		return
	}
	us, err := service.GetInfo(phone.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"user_info": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"user_info": us,
	})
}

func GetOInfo(c *gin.Context) {
	phone := c.PostForm("phone")
	oInfo, err := service.GetOInfo(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"user_info": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"user_info": oInfo,
	})
}

//GetRandCode 获取网站随机码
func GetRandCode(c *gin.Context) {
	phone := c.PostForm("phone")
	randCode, err := service.GetRandCode(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"rand_code": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"rand_code": randCode,
	})
}

func UpdatePwd(c *gin.Context) {
	phone := c.PostForm("phone")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	err := service.UpdatePwd(phone, oldPassword, newPassword)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
	})
}

func GetAnswer(c *gin.Context) {
	phone, res := c.Get("phone")
	if !res {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "user_name is not in the context",
			"answer": "",
		})
		return
	}
	answer, err := service.GetAnswer(phone.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"answer": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"answer": answer,
	})

}

func GetQuestion(c *gin.Context) {
	phone, res := c.Get("phone")
	if !res {
		c.JSON(200, gin.H{
			"status":   "1",
			"error":    "the phone is not in the gin of context,please check status of login!",
			"question": "",
		})
		return
	}
	question, err := service.GetQuestion(phone.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"status":   "0",
			"error":    fmt.Sprintf("%s", err),
			"question": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":   "1",
		"error":    "",
		"question": question,
	})
}

func GetToken(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user model.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"token":  "",
		})
		return
	}
	//校验用户名和密码是否正确
	rightPwd, err := dao.SelectUserPwd(user.Phone)
	if err == nil && rightPwd == tool.Encrypt(user.Password) {
		// 生成Token
		tokenString, err := tool.CreateToken(user.Phone)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "0",
				"error":  fmt.Sprintf("%s", err),
				"token":  "",
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "1",
			"error":  "",
			"token":  tokenString,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "0",
		"error":  fmt.Sprintf("%s", err),
		"token":  "",
	})
}

//GetErrToken 服务端返回错误token
func GetErrToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"token":  "this is a wrong token",
	})
}

func GetWODMvs(c *gin.Context) {
	label := c.Param("label")
	phone := c.GetString("phone")
	wodMvs, err := service.GetWODMvs(label, phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"movies": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"movies": wodMvs,
	})
}

func GetLComments(c *gin.Context) {
	phone := c.GetString("phone")
	lComments, err := service.GetLComments(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"status":        "0",
			"error":         fmt.Sprintf("%s", err),
			"long_comments": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":        "1",
		"error":         "",
		"long_comments": lComments,
	})
}
