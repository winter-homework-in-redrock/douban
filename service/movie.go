package service

import (
	"douban/dao"
	"douban/model"
	"douban/tool"
	"errors"
	"strconv"
	"time"
)

func GetMvInfo(id string) (model.Movie, error) {
	//转化为十进制int数据
	mvId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return model.Movie{}, err
	}
	movie, err := dao.SelectMvById(int(mvId))
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func GetHotMvs() ([]model.OfMovie, error) {
	hotMvs, err := dao.SelectMvsByHot()
	return hotMvs, err
}

func GetFutureMvs() ([]model.OfMovie, error) {
	futureMvs, err := dao.SelectMvsByFuture()
	return futureMvs, err
}

func CreateUDiscuss(fPhone, tPhone, dId, mvId, content string) error {
	//处理并组装数据
	id1, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return err
	}
	id2, err := strconv.ParseUint(dId, 10, 64)
	if err != nil {
		return err
	}
	var ud = model.UnderDiscuss{
		FromPhone: fPhone,
		ToPhone:   tPhone,
		Id:        int(id2),
		FromMvId:  int(id1),
		Content:   content,
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	err = dao.InsertUDiscuss(ud)
	return err
}

func DelDisOfMv(dId, phone string) error {
	id, err := strconv.ParseUint(dId, 10, 64)
	if err != nil {
		return err
	}
	err = dao.DelDisOfMv(int(id), phone)
	return err
}

func DelUDisOfMv(dId, phone string) error {
	id, err := strconv.ParseUint(dId, 10, 64)
	if err != nil {
		return err
	}
	err = dao.DelUDisOfMv(int(id), phone)
	return err
}

func CreateDiscuss(mvId string, phone interface{}, title, content string) error {
	//处理数据并校验
	if len(content) > 1000 {
		return errors.New("content is not >1000")
	}
	p := phone.(string)
	id, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return err
	}
	discuss := model.Discuss{
		FromPhone: p,
		FromMvId:  int(id),
		Title:     title,
		Content:   content,
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	err = dao.InsertDiscuss(discuss)
	return err
}

func GetDisOfMv(mvId string) ([]model.Discuss, error) {
	id, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return nil, err
	}
	disOfMv, err := dao.SelectDisOfMv(int(id))
	return disOfMv, err
}

func GetUDisOfMv(dId string) (map[int]model.UnderDiscuss, error) {
	id, err := strconv.ParseUint(dId, 10, 64)
	if err != nil {
		return nil, err
	}
	ulsOfMv, err := dao.GetUDisOfMv(int(id))
	return ulsOfMv, err
}

func GetMvsOfT(form, kind, place, age, special string) ([]model.OfMovie, error) {
	//未选参数，默认返回热门电影
	if form == "" && kind == "" && age == "" && place == "" && special == "" {
		mvsByHot, err := dao.SelectMvsByHot()
		if err != nil {
			return nil, err
		}
		return mvsByHot, nil
	}
	//todo 检验参数
	mvsOfT, err := dao.SelectMvsOfT(form, kind, place, age, special)
	return mvsOfT, err
}

func GetMvsOfR(typeName string) ([]model.OfMovie, error) {
	//校验参数，防止sql注入
	//中文正则
	err := tool.RegexChinese(typeName)
	if err != nil {
		return nil, err
	}
	mvsOfR, err := dao.SelectMvsOfR(typeName)
	return mvsOfR, err
}

func GetMvsOPerfs(search string) (model.OfStaff, []model.OfMovie, error) {
	//数据校验，约定只准输入中英文数字，不准输入特殊字符
	err := tool.RegexSearch(search)
	if err != nil {
		return model.OfStaff{}, nil, err
	}
	perf, ofMvs, err2 := dao.SelectMvsOfPerfs(search)
	return perf, ofMvs, err2
}
