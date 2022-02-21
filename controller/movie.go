package controller

import (
	"douban/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetMvInfo(c *gin.Context) {
	id := c.Param("mv_id")
	movie, err := service.GetMvInfo(id)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"movie":  "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"movie":  movie,
	})
}

func GetHotMvs(c *gin.Context) {
	hotMvs, err := service.GetHotMvs()
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
		"movies": hotMvs,
	})
}

func GetFutureMvs(c *gin.Context) {
	futureMvs, err := service.GetFutureMvs()
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
		"movies": futureMvs,
	})
}

func CreateUDiscuss(c *gin.Context) {
	//获取数据
	fPhone := c.GetString("phone")
	dId := c.Param("discuss_id")
	mvId := c.Param("mv_id")
	tPhone := c.PostForm("to_phone")
	content := c.PostForm("content")
	err := service.CreateUDiscuss(fPhone, tPhone, dId, mvId, content)
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

func DelDisOfMv(c *gin.Context) {
	dId := c.Param("discuss_id")
	phone := c.GetString("phone")
	err := service.DelDisOfMv(dId, phone)
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

func DelUDisOfMv(c *gin.Context) {
	dId := c.Param("discuss_id")
	phone := c.GetString("phone")
	err := service.DelUDisOfMv(dId, phone)
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

func CreateDiscuss(c *gin.Context) {
	//获取数据
	mvId := c.Param("mv_id")
	phone := c.GetString("phone")
	title := c.PostForm("title")
	content := c.PostForm("content")
	err := service.CreateDiscuss(mvId, phone, title, content)
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

func GetDisOfMv(c *gin.Context) {
	//获取数据
	mvId := c.Param("mv_id")
	disOfMv, err := service.GetDisOfMv(mvId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":    "0",
			"error":     fmt.Sprintf("%s", err),
			"discusses": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "1",
		"error":     "",
		"discusses": disOfMv,
	})
}

func GetUDisOfMv(c *gin.Context) {
	dId := c.Param("discuss_id")
	udOfMv, err := service.GetUDisOfMv(dId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":        "0",
			"error":         fmt.Sprintf("%s", err),
			"under_discuss": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":        "1",
		"error":         "",
		"under_discuss": udOfMv,
	})
}

func GetMvsOfT(c *gin.Context) {
	//电影分类标签，包括形式、类型、地区、年代、特色五个模块
	f := c.Query("form")
	k := c.Query("kind")
	p := c.Query("place")
	a := c.Query("age")
	s := c.Query("special")
	mvsOfT, err := service.GetMvsOfT(f, k, p, a, s)
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
		"movies": mvsOfT,
	})
}

func GetMvsOfR(c *gin.Context) {
	//获取分类名参数，默认为空
	typeName := c.Query("type_name")
	mvsOfR, err := service.GetMvsOfR(typeName)
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
		"movies": mvsOfR,
	})
}

func GetMvsOPerfs(c *gin.Context) {
	search := c.PostForm("search")
	perf, ofMvs, err := service.GetMvsOPerfs(search)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"staff":  "",
			"movies": "",
		})
		return
	}
	if perf.Id == "" {
		c.JSON(200, gin.H{
			"status": "1",
			"error":  "",
			"staff":  "",
			"movies": ofMvs,
		})
	} else {
		c.JSON(200, gin.H{
			"status": "1",
			"error":  "",
			"staff":  perf,
			"movies": ofMvs,
		})
	}
}
