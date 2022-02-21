package model

//Staff 影视演职员数据表结构映射
type Staff struct {
	Id            string `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Sex           string `json:"sex" form:"sex"`
	Avatar        string `json:"avatar" form:"avatar"`               //头像名
	Constellation string `json:"constellation" form:"constellation"` //星座
	Birthday      string `json:"birthday" form:"birthday"`           //出生日期
	Birthplace    string `json:"birthplace" form:"birthplace"`       //出生地
	Jobs          string `json:"jobs" form:"jobs"`                   //从事过的工作
	ACName        string `json:"ac_name" form:"ac_name"`             //中文别名
	AEName        string `json:"ae_name" form:"ae_name"`             //英文别名
	Family        string `json:"family" form:"family"`               //家庭成员
	Imdb          string `json:"imdb" form:"imdb"`                   //世界最大电影数据库编号
	Introduction  string `json:"introduction" form:"introduction"`   //介绍
}

//OfStaff 影职员部分数据
type OfStaff struct {
	Id     string `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Avatar string `json:"avatar" form:"avatar"` //头像名
	Jobs   string `json:"jobs" form:"jobs"`     //从事过的工作
}
