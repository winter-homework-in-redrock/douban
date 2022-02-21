package dao

import (
	"douban/global"
	"douban/model"
	"douban/tool"
	"errors"
	"fmt"
	"log"
)

func InsertWODMv(c model.ShortComment) error {
	//sql事务
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	switch c.MvStar {
	case 0:
		break
	case 1:
		sql1 := "update movie set mv_one_star_num=mv_one_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql1, c.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 2:
		sql1 := "update movie set mv_two_star_num=mv_two_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql1, c.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 3:
		sql1 := "update movie set mv_three_star_num=mv_three_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql1, c.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 4:
		sql1 := "update movie set mv_four_star_num=mv_four_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql1, c.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 5:
		sql1 := "update movie set mv_five_star_num=mv_five_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql1, c.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	default:
		return errors.New("the the mvStar is not in (0,1,2,3,4,5)")
	}

	sql2 := "insert into `user_movie` (mv_id,phone,label)values(?,?,?)"
	_, err = tx.Exec(sql2, c.FromMvId, c.FromPhone, c.WantOrWatched)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
	}
	//todo 先检测数据库是否存在已有评论。
	sql3 := "insert into `movie_short_comment` (from_phone,from_mv_id,want_or_watched,comment_mv_star,comment_tag,comment_content,comment_time,used_num)values (?,?,?,?,?,?,?,default)"
	_, err = tx.Exec(sql3, c.FromPhone, c.FromMvId, c.WantOrWatched, c.MvStar, c.Tag, c.Content, c.DateTime)
	if err != nil {
		err = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func DelScOfMv(mvId, cId int, phone string) error {
	var mvStar int
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	//查询评论评分
	sql0 := "select comment_mv_star from `movie_short_comment` where from_mv_id=? and from_phone=?"
	err = tx.QueryRow(sql0, mvId, phone).Scan(&mvStar)
	if err != nil {
		return err
	}
	//删除评分
	switch mvStar {
	case 0:
		break
	case 1:
		sql1 := "update movie set mv_one_star_num=mv_one_star_num-1 where mv_id=?"
		_, err = tx.Exec(sql1, mvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 2:
		sql1 := "update movie set mv_two_star_num=mv_two_star_num-1 where mv_id=?"
		_, err = tx.Exec(sql1, mvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 3:
		sql1 := "update movie set mv_three_star_num=mv_three_star_num-1 where mv_id=?"
		_, err = tx.Exec(sql1, mvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 4:
		sql1 := "update movie set mv_four_star_num=mv_four_star_num-1 where mv_id=?"
		_, err = tx.Exec(sql1, mvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 5:
		sql1 := "update movie set mv_five_star_num=mv_five_star_num-1 where mv_id=?"
		_, err = tx.Exec(sql1, mvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	default:
		return errors.New("the the mvStar is not in (0,1,2,3,4,5)")
	}
	//删除用户想看或看过的该电影
	sql2 := "delete from `user_movie` where mv_id=? and phone=?"
	_, err = tx.Exec(sql2, mvId, phone)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
	}
	//删除短评表
	sql3 := "delete from `movie_short_comment` where comment_id=?"
	_, err = tx.Exec(sql3, cId)
	if err != nil {
		err = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func SelectScsOfMv(mvId int) ([]model.ShortComment, error) {
	var sc = model.ShortComment{}
	var scs = make([]model.ShortComment, 0)
	//查询短评表数据
	sql1 := "select user_name,avatar from `user_side` where phone=?"
	sql2 := "select * from `movie_short_comment` where from_mv_id=?"
	rows, err := tool.DB.Query(sql2, mvId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(
			&sc.Id,
			&sc.FromPhone,
			&sc.FromMvId,
			&sc.WantOrWatched,
			&sc.MvStar,
			&sc.Tag,
			&sc.Content,
			&sc.DateTime,
			&sc.UsedNum,
		)
		if err != nil {
			return nil, err
		}
		//查询用户非隐私表数据
		err = tool.DB.QueryRow(sql1, sc.FromPhone).Scan(&sc.FromUserName, &sc.FromAvatar)
		if err != nil {
			return nil, err
		}
		//拼接图片路径
		sc.FromAvatar = global.UserAvatarPath+fmt.Sprintf("%s", sc.FromAvatar)
		scs = append(scs, sc)
	}
	return scs, nil
}

func InsertLcComment(lc model.LongComment) error {
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	//插入用户看过的电影
	sql1 := "insert into `user_movie` (mv_id,phone,label) values (?,?,default);"
	_, err = tx.Exec(sql1, lc.FromMvId, lc.FromPhone)
	if err != nil {
		tx.Rollback()
		return err
	}
	//插入影评数据
	sql2 := "insert into `movie_long_comment` (from_phone,from_mv_id,comment_mv_star,comment_title,comment_content,comment_time,used_num,unused_num) values (?,?,?,?,?,?,default,default);"
	_, err = tx.Exec(sql2, lc.FromPhone, lc.FromMvId, lc.MvStar, lc.Title, lc.Content, lc.DateTime)
	if err != nil {
		tx.Rollback()
		return err
	}
	//更新电影评分
	switch lc.MvStar {
	case 0:
		break
	case 1:
		sql3 := "update movie set mv_one_star_num=mv_one_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql3, lc.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 2:
		sql3 := "update movie set mv_two_star_num=mv_two_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql3, lc.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 3:
		sql3 := "update movie set mv_three_star_num=mv_three_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql3, lc.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 4:
		sql3 := "update movie set mv_four_star_num=mv_four_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql3, lc.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	case 5:
		sql3 := "update movie set mv_five_star_num=mv_five_star_num+1 where mv_id=?"
		_, err = tx.Exec(sql3, lc.FromMvId)
		if err != nil {
			tx.Rollback()
			return err
		}
	default:
		return errors.New("the the mvStar is not in (0,1,2,3,4,5)")
	}
	return tx.Commit()
}

func InsertULComment(Ulc model.UnderLongComment) error {
	sql := "insert into under_long_comment (comment_id,from_phone,to_phone,from_mv_id,comment_content,comment_time) values (?,?,?,?,?,?);"
	_, err := tool.DB.Exec(sql, Ulc.Id, Ulc.FromPhone, Ulc.ToPhone, Ulc.FromMvId, Ulc.Content, Ulc.DateTime)
	return err
}

func DelLsOfMv(cid int, phone string) error {
	tx, err := tool.DB.Begin()
	if err != nil {
		return err
	}
	//检查用户是否发表该影评
	var p string
	sql0 := "select from_phone from `movie_long_comment` where comment_id=?;"
	err = tx.QueryRow(sql0, cid).Scan(&p)
	if err != nil && p != phone {
		return errors.New("不能删除")
	}
	//删除影评
	sql1 := "delete from `movie_long_comment` where comment_id=? and from_phone =?"
	_, err = tx.Exec(sql1, cid, phone)
	if err != nil {
		tx.Rollback()
		return err
	}

	//删除影评下的回复
	sql2 := "delete from `under_long_comment` where comment_id=?;"
	_, err = tx.Exec(sql2, cid)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func DelULOfMv(phone string, cId int) error {
	sql := "update `under_long_comment` set comment_content=? where phone=? and comment_id=?"
	_, err := tool.DB.Exec(sql, "【已删除】", phone, cId)
	return err
}

func GetLcsOfMv(mvId int) ([]model.LongComment, error) {
	var lc = model.LongComment{}
	var lcs = make([]model.LongComment, 0)
	sql1 := "select * from `movie_long_comment` where from_mv_id=?"
	sql2 := "select user_name,avatar from `user_side` where phone=?"
	rows, err := tool.DB.Query(sql1, mvId)
	if err != nil {
		return nil, err
	}
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
		//查询头像名和用户名
		err = tool.DB.QueryRow(sql2, lc.FromPhone).Scan(&lc.FromUserName, &lc.FromAvatar)
		if err != nil {
			return nil, err
		}
		//拼接头像路径
		lc.FromAvatar = global.UserAvatarPath+fmt.Sprintf("%s", lc.FromAvatar)
		lcs = append(lcs, lc)
	}
	return lcs, nil
}

var level2 int

func GetUlsOfMv(cId int) (map[int]model.UnderLongComment, error) {
	var uLc = model.UnderLongComment{}
	var uLcs = make(map[int]model.UnderLongComment)
	var p string
	//查找顶级影评的发表者手机
	err := tool.DB.QueryRow("select from_phone from `movie_long_comment` where comment_id=?", cId).Scan(&p)
	if err != nil {
		log.Println("1,", err)
		return nil, err
	}
	wa, err := lTaoWa(p, uLc, uLcs)
	return wa, err
}
func lTaoWa(fromPhone string, uLc model.UnderLongComment, uLcs map[int]model.UnderLongComment) (map[int]model.UnderLongComment, error) {
	sql1 := "select user_name,avatar from `user_side` where phone=?"
	sql2 := "select * from `under_long_comment` where to_phone=?;"
	rows, err := tool.DB.Query(sql2, fromPhone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	log.Println("tao wa")
	for rows.Next() {
		log.Println("循环读取")
		//获取回应数据
		err = rows.Scan(
			&uLc.Id,
			&uLc.FromPhone,
			&uLc.ToPhone,
			&uLc.FromMvId,
			&uLc.Content,
			&uLc.DateTime,
		)
		err = tool.DB.QueryRow(sql1, uLc.FromPhone).Scan(&uLc.FromUserName, &uLc.FromAvatar)
		if err != nil {
			log.Println("3, ", err)
			return nil, err
		}
		//拼接头像路径
		uLc.FromAvatar = global.UserAvatarPath+fmt.Sprintf("%s", uLc.FromAvatar)
		uLcs[level2] = uLc
		level2++
		log.Println("ulc", uLc)
		//递归调用
		uLcs, err = lTaoWa(uLc.FromPhone, uLc, uLcs)
		if err != nil {
			log.Println("4, ", err)
			return nil, err
		}
	}
	return uLcs, nil
}
