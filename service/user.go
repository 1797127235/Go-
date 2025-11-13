package service

import (
	"blog-server/global"
	"blog-server/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//用户业务逻辑
func CreateUser(username,password string) error{
	//检查用户是否存在
	var count int64
	global.DB.Model(&model.User{}).Where("username = ?",username).Count(&count)
	if count > 0 {
		return fmt.Errorf("username already exists")
	}

	//加密密码
	hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	u := model.User{
		Username: username,
		Password: string(hash),
	}

	return global.DB.Create(&u).Error
}


//校验输入的密码是否与数据库中的加密密码一致
func CheckUser(username,password string) (*model.User, error){
	var u model.User
	if err := global.DB.Where("username = ?",username).First(&u).Error; err != nil{
		return nil,err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password)); err != nil{
		return nil,err
	}

	return &u,nil
}