package controller

import (
	"douban/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateWODMv(c *gin.Context) {
	//获取身份标识等数据
	mvId := c.Param("mv_id")
	phone, res := c.Get("phone")
	if res == false {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  "failed to get the 'phone' in the gin context,please check the JWTMiddleware in the application",
		})
		return
	}
	//获取短评内容、标签，默认为空
	content := c.DefaultPostForm("content", "")
	tag := c.DefaultPostForm("tag", "")
	//获取想看或看过数据  label:只能用0-->想看,1-->看过，默认为看过
	label := c.DefaultPostForm("label", "1")
	//获取评分，默认为0，表示不参与评分
	score := c.DefaultPostForm("score", "0")

	err := service.CreateWODMv(label, mvId, content, tag, score, phone)
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

func DelScOfMv(c *gin.Context) {
	mvId := c.Param("mv_id")
	cId := c.Param("comment_id")
	phone, _ := c.Get("phone")
	err := service.DelScOfMv(mvId, cId, phone)
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

func GetScsOfMv(c *gin.Context) {
	mvId := c.Param("mv_id")
	scsOfMv, err := service.GetScsOfMv(mvId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":         "0",
			"error":          fmt.Sprintf("%s", err),
			"short_comments": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":         "1",
		"error":          "",
		"short_comments": scsOfMv,
	})
}

func CreateLComment(c *gin.Context) {
	//获取数据
	mvId := c.Param("mv_id")
	phone := c.GetString("phone")
	title := c.PostForm("title")
	content := c.PostForm("content")
	score := c.DefaultPostForm("score", "0") //默认为零表示不参与评分
	err := service.CreateLComment(mvId, phone, title, content, score)
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

func CreateULComment(c *gin.Context) {
	//获取数据
	fPhone := c.GetString("phone")
	cId := c.Param("comment_id")
	mvId := c.Param("mv_id")
	tPhone := c.PostForm("to_phone")
	content := c.PostForm("content")

	err := service.CreateULComment(fPhone, tPhone, cId, mvId, content)
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

func DelLsOfMv(c *gin.Context) {
	cId := c.Param("comment_id")
	phone := c.GetString("phone")
	err := service.DelLsOfMv(cId, phone)
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

func DelULOfMv(c *gin.Context) {
	cId := c.Param("comment_id")
	phone := c.GetString("phone")
	err := service.DelULOfMv(cId, phone)
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

func GetLcsOfMv(c *gin.Context) {
	mvId := c.Param("mv_id")
	lcsOfMv, err := service.GetLcsOfMv(mvId)
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
		"long_comments": lcsOfMv,
	})
}

func GetUlsOfMv(c *gin.Context) {
	cId := c.Param("comment_id")
	ulsOfMv, err := service.GetUlsOfMv(cId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":              "0",
			"error":               fmt.Sprintf("%s", err),
			"under_long_comments": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":              "1",
		"error":               "",
		"under_long_comments": ulsOfMv,
	})
}
