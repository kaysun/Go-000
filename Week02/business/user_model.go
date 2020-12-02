package business

import (
	"fmt"
	"github.com/kaysun/Go-000/Week02/dao"
	"github.com/pkg/errors"
)

type UserModel struct {
	UserDao dao.UserDao
}

func(model UserModel) DemandUser (userID int)(dao.User, error) {
	user, err := model.UserDao.DemandUserByDB(userID)
	err = errors.WithMessage(err, fmt.Sprintf("demand user info by userID=%d fail", userID))
	return user, err
}