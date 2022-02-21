package service

import (
	"douban/dao"
	"douban/model"
	"strconv"
)

func GetStaffInfo(id string) (model.Staff, error) {

	staffId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return model.Staff{}, err
	}
	staffInfo, err := dao.SelectStaffInfo(int(staffId))
	return staffInfo, err
}

func GetStaffsOfMv(mvId string) ([]model.OfStaff, error) {
	//转化为十进制int数据
	id, err := strconv.ParseUint(mvId, 10, 64)
	if err != nil {
		return nil, err
	}
	staffsOfMv, err := dao.SelectStaffsOfMv(int(id))
	return staffsOfMv, err
}

func GetMvsOfStaff(sId string) ([]model.OfMovie, error) {
	//转化为十进制int数据
	id, err := strconv.ParseUint(sId, 10, 64)
	if err != nil {
		return nil, err
	}
	staffsOfMv, err := dao.SelectMvsOfStaff(int(id))
	return staffsOfMv, err
}
