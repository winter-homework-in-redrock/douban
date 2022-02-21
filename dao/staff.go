package dao

import (
	"douban/global"
	"douban/model"
	"douban/tool"
	"fmt"
	"log"
)

func SelectStaffInfo(id int) (model.Staff, error) {
	staff := model.Staff{}
	sql := "select * from staff where staff_id=?"
	err := tool.DB.QueryRow(sql, id).Scan(
		&staff.Id,
		&staff.Name,
		&staff.Sex,
		&staff.Avatar,
		&staff.Constellation,
		&staff.Birthday,
		&staff.Birthplace,
		&staff.Jobs,
		&staff.ACName,
		&staff.AEName,
		&staff.Family,
		&staff.Imdb,
		&staff.Introduction,
	)
	//拼接图片路径
	staff.Avatar = global.PerfPicturePath+fmt.Sprintf("%s", staff.Avatar)
	return staff, err
}

func SelectStaffsOfMv(mvId int) ([]model.OfStaff, error) {
	ids := make([]int, 0)
	var id int
	var staff model.OfStaff
	staffsOfMv := make([]model.OfStaff, 0)
	//从movie_staff数据表获取对应影视下的所有演职员id
	sql1 := "select staff_id from `movie_staff` where mv_id=?;"
	rows, err := tool.DB.Query(sql1, mvId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	//根据id遍历staff数据表获取演职员详细数据
	sql2 := "select staff_id,staff_name,staff_avatar,staff_jobs from `staff` where staff_id=?"
	for _, id := range ids {
		err = tool.DB.QueryRow(sql2, id).Scan(
			&staff.Id,
			&staff.Name,
			&staff.Avatar,
			&staff.Jobs,
		)
		if err != nil {
			return nil, err
		}
		//拼接路径
		staff.Avatar =global.PerfPicturePath+ fmt.Sprintf("%s", staff.Avatar)
		staffsOfMv = append(staffsOfMv, staff)
	}
	return staffsOfMv, nil
}

func SelectMvsOfStaff(sId int) ([]model.OfMovie, error) {
	ids := make([]int, 0)
	var id int
	var mv = model.OfMovie{}
	var mvs = make([]model.OfMovie, 0)
	//从movie_staff数据表获取对应影视下的所有演职员id
	sql1 := "select mv_id from `movie_staff` where staff_id =?;"
	rows, err := tool.DB.Query(sql1, sId)
	if err != nil {
		log.Println("1 ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Println("2 ", err)
			return nil, err
		}
		ids = append(ids, id)
	}
	log.Println(ids)
	//根据id遍历movie数据表获取电影数据
	var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum int
	sql2 := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num from `movie` where mv_id=?"
	for _, id := range ids {
		err = tool.DB.QueryRow(sql2, id).Scan(
			&mv.Id,
			&mv.Name,
			&mv.Picture,
			&mv.Director,
			&mv.LeadRole,
			&mv.ProduceWhere,
			&mv.ReleaseTime,
			&mv.Duration,
			&oneStarNum,
			&twoStarNum,
			&threeStarNum,
			&fourStarNum,
			&fiveStarNum,
		)
		if err != nil {
			log.Println("3 ", err)
			return nil, err
		}
		//防止除0错误
		if oneStarNum+twoStarNum+threeStarNum+fourStarNum+fiveStarNum == 0 {
			mv.Score = 0
		} else {
			//合并
			mv.Score = (oneStarNum + twoStarNum*2 + threeStarNum*3 + 4*fourStarNum + 5*fiveStarNum) / (oneStarNum + twoStarNum + threeStarNum + fourStarNum + fiveStarNum)
		}
		//拼接路径
		mv.Picture = global.MvPicturePath+fmt.Sprintf("%s", mv.Picture)
		mvs = append(mvs, mv)
	}
	return mvs, nil
}
