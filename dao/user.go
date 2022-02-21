package dao

import (
	"douban/global"
	"douban/model"
	"douban/tool"
	"fmt"
)

//InsertUser 同时插入隐私表、非隐私表数据
func InsertUser(user1 model.User, user2 model.UserSide) error {
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	insertStr1 := "insert into `user_info` (user_name,password,phone,question,answer)values(?,?,?,?,?)"
	_, err1 := tx.Exec(insertStr1, user1.UserName, user1.Password, user1.Phone, user1.Question, user1.Answer)
	insertStr2 := "insert into `user_side` (user_name,phone,avatar,user_introduction,user_sign,register_time)values(?,?,default ,default ,default,? )"
	_, err2 := tx.Exec(insertStr2, user2.UserName, user2.Phone, user2.RegisterTime)
	if err1 == nil && err2 == nil {
		err = tx.Commit()
		return err
	}
	err = tx.Rollback()
	return err2
}

func UpdateUserPwd(phone, userPwd string) error {
	updateStr := "update `user_info` set password=? where phone=?"
	_, err := tool.DB.Exec(updateStr, userPwd, phone)
	return err
}

func SelectUserPwd(phone string) (string, error) {
	var pwd string
	selectStr := "select password from `user_info` where phone =?"
	err := tool.DB.QueryRow(selectStr, phone).Scan(&pwd)
	return pwd, err
}

func SelectUserPhone(phone string) (string, error) {
	var un string
	selectStr := "select user_name from `user_info` where phone=?"
	err := tool.DB.QueryRow(selectStr, phone).Scan(&un)
	return un, err
}

func SelectUserAnswer(phone string) (string, error) {
	var answer string
	selectStr := "select answer from `user_info` where phone =?"
	err := tool.DB.QueryRow(selectStr, phone).Scan(&answer)
	return answer, err
}

func SelectUserQuestion(phone string) (string, error) {
	var question string
	selectStr := "select question from `user_info` where phone =?"
	err := tool.DB.QueryRow(selectStr, phone).Scan(&question)
	return question, err
}

func InsertUserIntroduction(phone, introduction string) error {
	str := "update `user_side` set user_introduction =? where phone=?"
	_, err := tool.DB.Exec(str, introduction, phone)
	return err
}

func InsertUserSign(phone, sign string) error {
	str := "update `user_side` set user_sign = ? where phone=?"
	_, err := tool.DB.Exec(str, sign, phone)
	return err
}

func InsertUserAvatar(phone, filename string) error {
	str := "update `user_side` set avatar = ? where phone=?"
	_, err := tool.DB.Exec(str, filename, phone)
	return err
}

func SelectUserSide(phone string) (model.UserSide, error) {
	us := model.UserSide{}
	str := "select * from `user_side` where phone=?"
	err := tool.DB.QueryRow(str, phone).Scan(
		&us.UserId,
		&us.UserName,
		&us.Phone,
		&us.Avatar,
		&us.UserIntroduction,
		&us.UserSign,
		&us.RegisterTime,
	)
	//拼接路径
	fileName := us.Avatar
	dst := global.UserAvatarPath+fmt.Sprintf("%s", fileName)
	us.Avatar = dst
	return us, err
}

func SelectOInfo(phone string) (model.UserSide, error) {
	us := model.UserSide{}
	str := "select * from `user_side` where phone=?"
	err := tool.DB.QueryRow(str, phone).Scan(
		&us.UserId,
		&us.UserName,
		&us.Phone,
		&us.Avatar,
		&us.UserIntroduction,
		&us.UserSign,
		&us.RegisterTime,
	)
	//拼接路径
	fileName := us.Avatar
	dst := global.UserAvatarPath+fmt.Sprintf("%s", fileName)
	us.Avatar = dst
	return us, err
}

func SelectWODMVs(label, phone string) ([]model.OfMovie, error) {
	var ids = make([]int, 0)
	var id int
	var ofMv = model.OfMovie{}
	var ofMvs = make([]model.OfMovie, 0)
	//查询user_movie表获取想看或看过的电影id
	sql1 := "select mv_id from `user_movie` where label=? and phone=?"
	rows, err := tool.DB.Query(sql1, label, phone)
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
	//依据id,查询影视具体数据
	var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum float32
	sql2 := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num from `movie` where mv_id=?"
	for _, id := range ids {
		err = tool.DB.QueryRow(sql2, id).Scan(
			&ofMv.Id,
			&ofMv.Name,
			&ofMv.Picture,
			&ofMv.Director,
			&ofMv.LeadRole,
			&ofMv.ProduceWhere,
			&ofMv.ReleaseTime,
			&ofMv.Duration,
			&oneStarNum,
			&twoStarNum,
			&threeStarNum,
			&fourStarNum,
			&fiveStarNum,
		)
		if err != nil {
			return nil, err
		}
		//四舍五入
		sc := (oneStarNum+twoStarNum*2+threeStarNum*3+4*fourStarNum+5*fiveStarNum)/(oneStarNum+twoStarNum+threeStarNum+fourStarNum+fiveStarNum) + 0.5
		ofMv.Score = int(sc)
		//拼接图片路径
		ofMv.Picture = global.MvPicturePath+fmt.Sprintf("%s", ofMv.Picture)
		ofMvs = append(ofMvs, ofMv)
	}
	return ofMvs, nil
}

func SelectLComments(phone string) ([]model.LongComment, error) {
	var lc = model.LongComment{}
	var lcs = make([]model.LongComment, 0)
	sql0 := "select user_name,avatar from `user_side` where phone= ?;"
	sql1 := "select * from `movie_long_comment` where from_phone = ?;"
	rows, err := tool.DB.Query(sql1, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&lc.Id,
			&lc.FromPhone,
			&lc.FromMvId,
			&lc.MvStar,
			&lc.Title,
			&lc.Content,
			&lc.DateTime,
			&lc.UsedNum,
			&lc.UnusedNum,
		)
		if err != nil {
			return nil, err
		}
		err = tool.DB.QueryRow(sql0, lc.FromPhone).Scan(&lc.FromUserName, &lc.FromAvatar)
		if err != nil {
			return nil, err
		}
		//拼接图片路径
		lc.FromAvatar = global.UserAvatarPath+fmt.Sprintf("%s", lc.FromAvatar)

		lcs = append(lcs, lc)
	}
	return lcs, nil
}
