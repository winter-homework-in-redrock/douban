package controller

import (
	"douban/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetStaffInfo(c *gin.Context) {
	id := c.Param("staff_id")
	staffInfo, err := service.GetStaffInfo(id)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"staff":  "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "0",
		"error":  "",
		"staff":  staffInfo,
	})

}

func GetStaffsOfMv(c *gin.Context) {
	mvId := c.Param("mv_id")
	staffsOfMv, err := service.GetStaffsOfMv(mvId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"staffs": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"staffs": staffsOfMv,
	})
}

func GetMvsOfStaff(c *gin.Context) {
	sId := c.Param("staff_id")
	staffsOfMv, err := service.GetMvsOfStaff(sId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "0",
			"error":  fmt.Sprintf("%s", err),
			"staffs": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "1",
		"error":  "",
		"staffs": staffsOfMv,
	})
}
