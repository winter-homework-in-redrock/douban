package dao

import (
	"database/sql"
	"douban/global"
	"douban/model"
	"douban/tool"
	"errors"
	"fmt"
	"log"
	"strings"
)

func SelectMvById(id int) (model.Movie, error) {
	mv := model.Movie{}
	sqlStr := "select * from `movie` where `mv_id` = ?"
	err := tool.DB.QueryRow(sqlStr, id).Scan(
		&mv.Id,
		&mv.Form,
		&mv.Name,
		&mv.Picture,
		&mv.Director,
		&mv.Writer,
		&mv.LeadRole,
		&mv.Type,
		&mv.Special,
		&mv.ProduceWhere,
		&mv.Language,
		&mv.ReleaseTime,
		&mv.Duration,
		&mv.AName,
		&mv.Imdb,
		&mv.PlotIntroduction,
		&mv.OneStarNum,
		&mv.TwoStarNum,
		&mv.ThreeStarNum,
		&mv.FourStarNum,
		&mv.FiveStarNum,
	)
	//拼接路径
	mv.Picture = global.MvPicturePath + fmt.Sprintf("%s", mv.Picture)
	return mv, err
}

func SelectMvsByHot() ([]model.OfMovie, error) {
	sql := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num from `movie` where mv_form ='电影' and DATEDIFF(NOW(),mv_release_time)<30"
	rows, err := tool.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var mv = model.OfMovie{}
	var mvs = make([]model.OfMovie, 0)
	var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum int
	for rows.Next() {
		err = rows.Scan(
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
		mv.Picture = global.MvPicturePath + fmt.Sprintf("%s", mv.Picture)
		mvs = append(mvs, mv)
	}
	return mvs, err
}

func SelectMvsByFuture() ([]model.OfMovie, error) {
	sql := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num from `movie` where mv_form ='电影' and DATEDIFF(NOW(),mv_release_time)<0"
	rows, err := tool.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var mv = model.OfMovie{}
	var mvs = make([]model.OfMovie, 0)
	var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum int
	for rows.Next() {
		err = rows.Scan(
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
		mv.Picture = global.MvPicturePath + fmt.Sprintf("%s", mv.Picture)
		mvs = append(mvs, mv)
	}
	return mvs, err
}

func InsertUDiscuss(ud model.UnderDiscuss) error {
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	//检查mv_id是否存在
	var name string
	sql0 := "select mv_name from `movie` where mv_id=?;"
	err = tx.QueryRow(sql0, ud.FromMvId).Scan(&name)
	if err != nil {
		return err
	}
	//检查discuss_id是否存在
	var c string
	sql1 := "select discuss_content from `movie_discuss` where discuss_id=?;"
	err = tx.QueryRow(sql1, ud.Id).Scan(&c)
	if err != nil {
		return err
	}
	//更新最后一条评论的标记
	sql2 := "update `discuss_under` set label=1 where from_discuss_id=?"
	_, err = tx.Exec(sql2, ud.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	//插入评论
	sql3 := "insert into `discuss_under` (from_discuss_id,from_phone,from_mv_id,to_phone,used_num,under_content,under_time,label) values (?,?,?,?,default,?,?,default);"
	_, err = tx.Exec(sql3, ud.Id, ud.FromPhone, ud.FromMvId, ud.ToPhone, ud.Content, ud.DateTime)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func DelDisOfMv(dId int, phone string) error {
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	//检查用户是否发表该影评
	var p string
	sql0 := "select from_phone from `movie_discuss` where discuss_id=?;"
	err = tx.QueryRow(sql0, dId).Scan(&p)
	if err != nil && p != phone {
		return errors.New("不能删除")
	}
	//删除讨论
	sql1 := "delete from `movie_discuss` where discuss_id=?"
	_, err = tx.Exec(sql1, dId)
	if err != nil {
		tx.Rollback()
		return err
	}
	//删除讨论下的回复
	sql2 := "delete from `discuss_under` where from_discuss_id=?;"
	_, err = tx.Exec(sql2, dId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func DelUDisOfMv(dId int, phone string) error {
	sql := "update `discuss_under` set under_content=? where from_phone=? and from_discuss_id=?"
	_, err := tool.DB.Exec(sql, "【已删除】", phone, dId)
	return err
}

func InsertDiscuss(discuss model.Discuss) error {
	//检查mv_id是否存在
	var name string
	sql0 := "select mv_name from `movie` where mv_id=?;"
	err := tool.DB.QueryRow(sql0, discuss.FromMvId).Scan(&name)
	if err != nil {
		return err
	}
	sql1 := "insert into `movie_discuss` (from_phone,from_mv_id,discuss_title,discuss_content,discuss_num,discuss_time) values (?,?,?,?,default,?);"
	_, err = tool.DB.Exec(sql1, discuss.FromPhone, discuss.FromMvId, discuss.Title, discuss.Content, discuss.DateTime)
	return err
}

func SelectDisOfMv(mvId int) ([]model.Discuss, error) {
	var discuss = model.Discuss{}
	var ds = make([]model.Discuss, 0)
	sql1 := "select user_name,avatar from `user_side` where phone=?"
	sql2 := "select * from `movie_discuss` where from_mv_id=?;"
	rows, err := tool.DB.Query(sql2, mvId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&discuss.Id,
			&discuss.FromPhone,
			&discuss.FromMvId,
			&discuss.Title,
			&discuss.Content,
			&discuss.DiscussNum,
			&discuss.DateTime,
		)
		if err != nil {
			return nil, err
		}
		//查询头像和用户名
		err = tool.DB.QueryRow(sql1, discuss.FromPhone).Scan(&discuss.FromUserName, &discuss.FromAvatar)
		if err != nil {
			return nil, err
		}
		//拼接头像路径
		discuss.FromAvatar = global.UserAvatarPath + fmt.Sprintf("%s", discuss.FromAvatar)
		//追加切片
		ds = append(ds, discuss)
	}
	return ds, nil
}

var level int

//GetUDisOfMv todo 套娃BUG
func GetUDisOfMv(dId int) (map[int]model.UnderDiscuss, error) {
	var p string
	//查找顶级讨论的发表者手机
	err := tool.DB.QueryRow("select from_phone from `movie_discuss` where discuss_id=?", dId).Scan(&p)
	if err != nil {
		return nil, err
	}
	wa, err := dTaoWa(p)
	return wa, err
}
func dTaoWa(fromPhone string) (map[int]model.UnderDiscuss, error) {
	var ud = model.UnderDiscuss{}
	var mapUd = make(map[int]model.UnderDiscuss)
	sql1 := "select user_name,avatar from `user_side` where phone=?;"
	sql2 := "select * from `discuss_under` where to_phone=?;"
	rows, err := tool.DB.Query(sql2, fromPhone)
	if err != nil {
		return mapUd, err
	}
	defer rows.Close()
	for rows.Next() {
		//获取回应数据
		err = rows.Scan(
			&ud.Id,
			&ud.FromPhone,
			&ud.FromMvId,
			&ud.ToPhone,
			&ud.UsedNum,
			&ud.Content,
			&ud.DateTime,
			&ud.Label,
		)
		if err != nil {
			return mapUd, err
		}
		err = tool.DB.QueryRow(sql1, ud.FromPhone).Scan(&ud.FromUserName, &ud.FromAvatar)
		if err != nil {
			return mapUd, err
		}
		//拼接头像路径
		ud.FromAvatar = global.UserAvatarPath + fmt.Sprintf("%s", ud.FromAvatar)
		//保存数据
		mapUd[level] = ud
		level++
		log.Println(level, ud)
		//递归结束条件
		if ud.Label == "1" {
			continue
		} else if ud.Label == "0" {
			break
		}
		//递归调用
		mapUd, err = dTaoWa(ud.FromPhone)
		if err != nil {
			return mapUd, err
		}
	}
	return mapUd, nil
}

func SelectMvsOfT(form, kind, place, age, special string) ([]model.OfMovie, error) {
	//想不出来了，只有上shi山代码了 ψ(._. )>
	//占位符数量
	var a, c int
	//拼接sql语句
	sql1 := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration from `movie` where "
	if form != "" {
		sql1 = sql1 + " mv_form=? and"
		a++
	}
	if kind != "" {
		sql1 = sql1 + " mv_type like '%" + kind + "%' and"
	}
	if place != "" {
		sql1 = sql1 + " mv_produce_where like '%" + place + "%'and"
	}
	if age != "" {
		sql1 = sql1 + " year(mv_release_time) = ? and"
		c++
	}
	if special != "" {
		sql1 = sql1 + " mv_special like '%" + special + "%' and"
	}
	//去掉尾部and
	sql1 = strings.TrimSuffix(sql1, "and")
	//确定占位符数量
	if a == 0 && c == 0 {
		rows, err := tool.DB.Query(sql1)
		if err != nil {
			return nil, err
		}
		ofMvs, err := shiShan(rows)
		if err != nil {
			return nil, err
		}
		return ofMvs, nil
	} else if a == 0 && c == 1 {
		rows, err := tool.DB.Query(sql1, age)
		if err != nil {
			return nil, err
		}
		ofMvs, err := shiShan(rows)
		if err != nil {
			return nil, err
		}
		return ofMvs, nil
	} else if a == 1 && c == 0 {
		rows, err := tool.DB.Query(sql1, form)
		if err != nil {
			return nil, err
		}
		ofMvs, err := shiShan(rows)
		if err != nil {
			return nil, err
		}
		return ofMvs, nil
	} else if a == 1 && c == 1 {
		rows, err := tool.DB.Query(sql1, form, age)
		if err != nil {
			return nil, err
		}
		ofMvs, err := shiShan(rows)
		if err != nil {
			return nil, err
		}
		return ofMvs, nil
	}
	return nil, errors.New("排行榜电影查询失败")
}

func shiShan(rows *sql.Rows) ([]model.OfMovie, error) {
	var ofMv = model.OfMovie{}
	var ofMvs = make([]model.OfMovie, 0)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&ofMv.Id,
			&ofMv.Name,
			&ofMv.Picture,
			&ofMv.Director,
			&ofMv.LeadRole,
			&ofMv.ProduceWhere,
			&ofMv.ReleaseTime,
			&ofMv.Duration,
		)
		if err != nil {
			return nil, err
		}
		ofMv.Picture = global.MvPicturePath + fmt.Sprintf("%s", ofMv.Picture)
		ofMvs = append(ofMvs, ofMv)
	}
	return ofMvs, nil
}

func SelectMvsOfR(typeName string) ([]model.OfMovie, error) {
	var ofMv = model.OfMovie{}
	var ofMvs = make([]model.OfMovie, 0)
	var sc []uint8
	//之所以这么绕，是因为报错：converting driver.Value type []uint8 (\"5.0000\") to a int: invalid syntax,好像mysql的ifnull返回的数据类型是[]uint8
	//模糊+多列排序查询
	sql1 := "SELECT mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration,mv_one_star_num,mv_two_star_num,mv_three_star_num,mv_four_star_num,mv_five_star_num,ifnull((mv_one_star_num+mv_two_star_num*2+mv_three_star_num*3+mv_four_star_num*4+mv_five_star_num*5)/(mv_one_star_num+mv_two_star_num+mv_three_star_num+mv_four_star_num+mv_five_star_num),0 )" +
		" AS score FROM `movie` WHERE mv_type LIKE " + "'%" + typeName + "%'" + " ORDER BY score DESC"
	rows, err := tool.DB.Query(sql1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var oneStarNum, twoStarNum, threeStarNum, fourStarNum, fiveStarNum int
	for rows.Next() {
		err = rows.Scan(
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
			&sc,
		)
		if err != nil {
			return nil, err
		}
		//防止除0错误
		if oneStarNum+twoStarNum+threeStarNum+fourStarNum+fiveStarNum == 0 {
			ofMv.Score = 0
		} else {
			//合并
			ofMv.Score = (oneStarNum + twoStarNum*2 + threeStarNum*3 + 4*fourStarNum + 5*fiveStarNum) / (oneStarNum + twoStarNum + threeStarNum + fourStarNum + fiveStarNum)
		}

		ofMv.Picture = global.MvPicturePath + fmt.Sprintf("%s", ofMv.Picture)
		ofMvs = append(ofMvs, ofMv)
	}
	return ofMvs, nil
}

func SelectMvsOfPerfs(search string) (model.OfStaff, []model.OfMovie, error) {
	var ofStaff = model.OfStaff{}
	var ofMv = model.OfMovie{}
	var ofMvs = make([]model.OfMovie, 0)
	//查询演职员信息
	sql1 := "select staff_id,staff_name,staff_avatar,staff_jobs from `staff` where staff_name like '%" + search + "%' limit 1"
	err := tool.DB.QueryRow(sql1).Scan(
		&ofStaff.Id,
		&ofStaff.Name,
		&ofStaff.Avatar,
		&ofStaff.Jobs,
	)

	//拼接路径
	ofStaff.Avatar = global.PerfPicturePath + fmt.Sprintf("%s", ofStaff.Avatar)
	//查询影视信息
	sql2 := "select mv_id,mv_name,mv_picture,mv_director,mv_lead_role,mv_produce_where,mv_release_time,mv_duration from `movie` where mv_name like '%" + search + "%'"
	rows, err := tool.DB.Query(sql2)
	if err != nil {
		return ofStaff, ofMvs, err
	}
	for rows.Next() {
		err = rows.Scan(
			&ofMv.Id,
			&ofMv.Name,
			&ofMv.Picture,
			&ofMv.Director,
			&ofMv.LeadRole,
			&ofMv.ProduceWhere,
			&ofMv.ReleaseTime,
			&ofMv.Duration,
		)
		if err != nil {
			return ofStaff, ofMvs, err
		}
		//拼接路径
		ofMv.Picture = global.MvPicturePath + fmt.Sprintf("%s", ofMv.Picture)
		ofMvs = append(ofMvs, ofMv)
	}
	return ofStaff, ofMvs, nil
}
