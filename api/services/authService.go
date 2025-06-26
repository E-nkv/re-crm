package services

import (
	"context"
	"log"
	"re-crm/dtos"
	"re-crm/errs"
	"re-crm/repositories"

	"re-crm/utils"
)

type AuthService struct {
	UserRepo       repositories.UserRepo
	BlacklistedIds utils.ThreadSafeSet[uint64]
}

func NewAuthService(repo repositories.UserRepo) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (authSvc *AuthService) Login(ctx context.Context, dto dtos.LoginDTO) (string, string, error) {
	user, err := authSvc.UserRepo.GetByNickPass(ctx, dto)
	if err != nil {
		log.Println("dto:", dto)
		log.Println("😡 err getting user by nickpass: ", err.Error())
		return "", "", err
	}

	isBlacklisted, err := authSvc.IsUserBlacklisted(ctx, user.ID)
	if err != nil {
		return "", "", errs.Internal
	}
	if isBlacklisted {
		return "", "", errs.NotAllowed
	}
	token, err := utils.GenerateJWT(user.ID, user.Role)
	return token, user.Role, err
}

// private methods
func (authSvc *AuthService) IsUserBlacklisted(ctx context.Context, userID uint64) (bool, error) {
	return authSvc.BlacklistedIds.Contains(userID), nil
}
