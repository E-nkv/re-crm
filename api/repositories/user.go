package repositories

import (
	"context"
	"errors"
	"re-crm/dtos"
	"re-crm/entities"
	"re-crm/errs"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetByNickPass(context.Context, dtos.LoginDTO) (*entities.User, error)
}

type UserRepoPg struct {
	UserRepo
	db *gorm.DB
}

func NewUserRepoPg(db *gorm.DB) *UserRepoPg {
	return &UserRepoPg{
		db: db,
	}
}
func (ur *UserRepoPg) GetByNickPass(ctx context.Context, dto dtos.LoginDTO) (*entities.User, error) {
	var u entities.User
	res := ur.db.Where("nick = ?", dto.Nick).First(&u)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errs.NotFound
		}
		return nil, res.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Pass), []byte(dto.Pass)); err != nil {
		return nil, errs.InvalidCreds
	}

	return &u, nil
}
