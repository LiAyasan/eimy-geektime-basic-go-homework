package service

import (
	"context"
	"eimy-geektime-basic-go-homework/webook/internal/domain"
	"eimy-geektime-basic-go-homework/webook/internal/repository"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// ErrUserDuplicateEmail 每一层都有自己的Error，不出现跨层耦合
var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 你要考虑加密放在哪里（公共方法？）
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// 然后入库
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	// 先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// DEBUG log
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserService) Edit(ctx context.Context, userId int64, nickname, birthday, details string) error {
	return svc.repo.UpdateById(ctx, userId, nickname, birthday, details)
}

func (svc *UserService) Profile(ctx context.Context, userId int64) (string, error) {
	user, err := svc.repo.FindById(ctx, userId)
	if err != nil {
		return "", err
	}
	userJson, _ := json.Marshal(user)
	return string(userJson), nil
}
