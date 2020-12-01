package dao

import (
	"github.com/pkg/errors"
)

var ErrNoRows  = errors.New("sql.ErrNoRows")

type User struct {
	UserID   int
	UserName string
}

type UserDao struct {
}

func (dao UserDao) DemandUserByDB(userID int)(User, error) {
	var user User
	err := errors.Wrap(ErrNoRows, "db中没有找到数据")
	err = errors.WithStack(ErrNoRows)
	return user, err
}
