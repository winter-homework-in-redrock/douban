package service

import (
	"douban/dao"
	"douban/model"
	"errors"
	"strconv"
	"time"
)

func CreateWODMv(label, mvId, content, tag, score string, phone interface{}) error {
	if len(tag) > 50 {
		return errors.New("the tag length is not > 350")
	}
	if len(content) > 350 {
		return errors.New("the content length is not > 350")
	}
	p := phone.(string)
	id, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return err
	}
	s, err := strconv.ParseUint(score, 10, 64)
	if err != nil {
		return err
	}
	if label != "0" && label != "1" {
		return errors.New("the label must be '0' or '1'")
	}
	//组装数据
	sc := model.ShortComment{
		FromPhone:     p,
		FromMvId:      int(id),
		WantOrWatched: label,
		MvStar:        int(s),
		Tag:           tag,
		Content:       content,
		DateTime:      time.Now().Format("2006-01-02 15:04:05"),
	}
	err = dao.InsertWODMv(sc)
	return err
}

func DelScOfMv(mvId, cId string, phone interface{}) error {
	id1, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return err
	}
	id2, err := strconv.ParseUint(cId, 10, 64)
	if err != nil {
		return err
	}
	p := phone.(string)
	err = dao.DelScOfMv(int(id1), int(id2), p)
	return err
}

func GetScsOfMv(mvId string) ([]model.ShortComment, error) {
	id1, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return nil, err
	}
	scsOfMv, err := dao.SelectScsOfMv(int(id1))
	return scsOfMv, err
}

func CreateLComment(mvId string, phone interface{}, title, content, score string) error {
	//处理数据并校验
	p := phone.(string)
	id, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return err
	}
	s, err := strconv.ParseUint(score, 10, 64)
	if err != nil {
		return err
	}

	lc := model.LongComment{
		FromPhone: p,
		FromMvId:  int(id),
		MvStar:    int(s),
		Title:     title,
		Content:   content,
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	err = dao.InsertLcComment(lc)
	return err
}

func CreateULComment(fPhone, tPhone, cId, mvId, content string) error {
	//处理并组装数据
	id1, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return err
	}
	id2, err := strconv.ParseUint(cId, 10, 64)
	if err != nil {
		return err
	}
	var Ulc = model.UnderLongComment{
		FromPhone: fPhone,
		ToPhone:   tPhone,
		Id:        int(id2),
		FromMvId:  int(id1),
		Content:   content,
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	err = dao.InsertULComment(Ulc)
	return err
}

func DelLsOfMv(cId, phone string) error {
	id, err := strconv.ParseUint(cId, 10, 64)
	if err != nil {
		return err
	}
	err = dao.DelLsOfMv(int(id), phone)
	return err
}

func DelULOfMv(cId string, phone interface{}) error {
	p := phone.(string)
	id, err := strconv.ParseUint(cId, 10, 64)
	if err != nil {
		return err
	}
	err = dao.DelULOfMv(p, int(id))
	return err
}
func GetLcsOfMv(mvId string) ([]model.LongComment, error) {
	id1, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return nil, err
	}
	lcsOfMv, err := dao.GetLcsOfMv(int(id1))
	if err != nil {
		return nil, err
	}
	return lcsOfMv, nil
}

func GetUlsOfMv(cId string) (map[int]model.UnderLongComment, error) {
	id, err := strconv.ParseUint(cId, 10, 64)
	if err != nil {
		return nil, err
	}
	ulsOfMv, err := dao.GetUlsOfMv(int(id))
	return ulsOfMv, err
}
