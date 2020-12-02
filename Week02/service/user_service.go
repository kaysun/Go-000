package service

import (
	"errors"
	"github.com/kaysun/Go-000/Week02/dao"

	"fmt"
	"github.com/kaysun/Go-000/Week02/business"
)

type UserService struct {
	UserModel business.UserModel
}

func(service UserService) DemandUser (UserID int) {
	service.UserModel = business.UserModel{}
	user, err := service.UserModel.DemandUser(UserID)
	flag := errors.Is(err, dao.ErrNoRows)
	fmt.Println(user, err)
	fmt.Println(fmt.Sprintf("err=%+v",err))
	fmt.Println(fmt.Sprintf("是否是ErrNoRows类型？flag=%v", flag))
}