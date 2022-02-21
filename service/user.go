package service

import (
	"douban/dao"
	"douban/model"
	"douban/tool"
	"errors"
	"fmt"
	"log"
	"time"
)

func Register(user model.User) error {
	//用户信息校验格式正则校验
	err := tool.RegexPhone(user.Phone)
	if err != nil {
		return err
	}
	err = tool.RegexUserNameAndPwd(user.UserName)
	if err != nil {
		return err
	}
	err = tool.RegexUserNameAndPwd(user.Password)
	if err != nil {
		return err
	}
	//检测随机码
	rand, err := tool.RedisGetExp(user.Phone)
	log.Println(rand, err)
	if err != nil {
		return errors.New("the rand code is out of date or you should put into right rand code")
	}
	//手机号存在的话，直接登录
	phone, err := dao.SelectUserPhone(user.Phone)
	if err == nil && user.Phone == phone {
		return nil
	}
	//密码加密
	user.Password = tool.Encrypt(user.Password)
	//用户名为空，获取随机用户名GoString implements fmt.GoStringer and formats t to be printed in Go source code
	if user.UserName == "" {
		user.UserName = tool.RandString()
	}
	us := model.UserSide{UserName: user.UserName, Phone: user.Phone, RegisterTime: time.Now().Format("2006-01-02")}
	//添加用户隐私数据以及非隐私数据表
	err = dao.InsertUser(user, us)
	fmt.Println(err)
	return err
}

func CreateIntroduction(phone, introduction string) error {
	//校验数据格式
	if len(introduction) > 1000 {
		return errors.New("the introduction is too long, please to write <1000 bytes")
	}
	err := dao.InsertUserIntroduction(phone, introduction)
	return err
}

func CreateSign(phone, sign string) error {
	//校验数据格式
	if len(sign) > 200 {
		return errors.New("the sign is too long, please to write <200 bytes")
	}
	err := dao.InsertUserSign(phone, sign)
	return err
}

func UploadUserAvatar(phone, filename string) error {
	//校验数据格式
	if len(filename) > 200 {
		return errors.New("the filename is too long, please to write <200 bytes")
	}
	//todo 删除原有头像文件
	err := dao.InsertUserAvatar(phone, filename)
	return err
}

func GetInfo(phone string) (model.UserSide, error) {
	us, err := dao.SelectUserSide(phone)
	if err != nil {
		return us, err
	}

	return us, err
}

func GetOInfo(phone string) (model.UserSide, error) {
	oInfo, err := dao.SelectOInfo(phone)
	return oInfo, err
}

func GetRandCode(phone string) (int, error) {
	err := tool.RegexPhone(phone)
	if err != nil {
		return -1, err
	}
	// 获取网站随机码充当验证码
	randNum, err := tool.RandNum()
	if err != nil {
		return randNum, err
	}
	err = tool.RedisSetExp(randNum,phone)
	return randNum, err
}

func LoginByPwd(phone, password string) error {
	userPwd, err := dao.SelectUserPwd(phone)
	if err != nil {
		return err
	}
	//加密算法
	password = tool.Encrypt(password)
	if userPwd != password {
		return errors.New("the user`s password is wrong")
	}
	return nil
}

func UpdatePwd(phone, oldPassword, newPassword string) error {
	//校验密码格式
	err := tool.RegexUserNameAndPwd(newPassword)
	if err != nil {
		return err
	}
	err = tool.RegexUserNameAndPwd(phone)
	if err != nil {
		return err
	}
	//验证密码是否正确
	pwd, err := dao.SelectUserPwd(phone)
	if err != nil {
		return err
	}
	//加密算法
	oldPassword = tool.Encrypt(oldPassword)
	log.Println("", oldPassword, "\t", pwd)
	if pwd != oldPassword {
		return errors.New("the password is wrong")
	}
	//加密算法
	newPassword = tool.Encrypt(newPassword)
	log.Println(newPassword)
	err = dao.UpdateUserPwd(phone, newPassword)
	return err
}

func GetAnswer(phone string) (string, error) {
	answer, err := dao.SelectUserAnswer(phone)
	if err != nil {
		return "", err
	}
	return answer, err
}

func GetQuestion(phone string) (string, error) {
	question, err := dao.SelectUserQuestion(phone)
	if err != nil {
		return "", err
	}
	return question, err
}

func GetWODMvs(label string, phone string) ([]model.OfMovie, error) {
	if label != "0" && label != "1" {
		return nil, errors.New("the label must be '0' or '1'")
	}
	wodMVs, err := dao.SelectWODMVs(label, phone)
	return wodMVs, err
}

func GetLComments(phone string) ([]model.LongComment, error) {
	lComments, err := dao.SelectLComments(phone)
	return lComments, err
}
