package model

import (
	"github.com/dgrijalva/jwt-go"
)

//User 保存注册用户隐私数据
type User struct {
	UserID   string `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`       //用户绑定的手机号码（唯一）
	Question string `json:"question" form:"question"` //忘记密码时回答的问题
	Answer   string `json:"answer" form:"answer"`     //忘记密码时的问题答案
}

//UserSide 保存注册用户非隐私数据
type UserSide struct {
	UserId           string `json:"user_id" form:"user_id"`                     //为对应用户隐私表主键ID
	Phone            string `json:"phone" form:"phone"`                         //当前登录用户的手机号
	Avatar           string `json:"avatar" form:"avatar"`                       //用户头像绝对路径
	UserName         string `json:"user_name" form:"user_name"`                 //用户的昵称,依据隐私表数据
	UserIntroduction string `json:"user_introduction" form:"user_introduction"` //用户的自我介绍
	UserSign         string `json:"user_sign" form:"user_sign"`                 //用户的签名档
	RegisterTime     string `json:"register_time" form:"register_time"`         //用户的注册时间

}

//UserInfo 绑定数据，仅方便用于获取账户密码
type UserInfo struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

//Claims 用于获取token
type Claims struct {
	Phone              string `json:"phone" form:"phone"` //自定义字段
	jwt.StandardClaims        //官方字段
}
