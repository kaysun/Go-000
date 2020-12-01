package service

import (
	"errors"
	"github.com/kaysun/Go-000/Week02/dao"

	"fmt"
	"github.com/kaysun/Go-000/Week02/model"
)

type UserService struct {
	UserModel model.UserModel
}

func(service UserService) DemandUser (UserID int) {
	service.UserModel = model.UserModel{}
	user, err := service.UserModel.DemandUser(UserID)
	flag := errors.Is(err, dao.ErrNoRows)
	fmt.Println(user, err)
	fmt.Println(fmt.Sprintf("err=%+v",err))
	fmt.Println(fmt.Sprintf("是否是ErrNoRows类型？flag=%v", flag))
}