package models

import (
	"DataCertProject01/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id int   `form:"id"`
	Phone string   `form:"phone"`
	Password string   `form:"password"`
}

func (u User) SaveUser()(int64 , error) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(bytes)
	row,err :=db_mysql.Db.Exec("insert into user(phone,password)"+"values(?,?)",u.Phone,u.Password)
	if err != nil {
		return -1, err
	}
	id,err :=row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (u User) QueryUser()(*User,error) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	fmt.Println("用户电话",u.Phone)
	bytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(bytes)
	fmt.Println("用户密码",u.Password)
	row :=db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone,u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}