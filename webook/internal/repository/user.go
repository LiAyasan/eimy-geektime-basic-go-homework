package repository

import (
	"context"
	"eimy-geektime-basic-go-homework/webook/internal/domain"
	"eimy-geektime-basic-go-homework/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	// 在 repo 这一层只有创建的概念
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// 先从 cache 找
	// 再从 dao 里面找
	// 找到了回写 cache
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.UserExtend, error) {
	u, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.UserExtend{}, err
	}
	return domain.UserExtend{
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		Details:  u.Details,
	}, nil
}

func (r *UserRepository) UpdateById(ctx context.Context, id int64, nickname, birthday, details string) error {
	// 补充空字段
	u, err := r.dao.FindById(ctx, id)
	if err != nil {
		return err
	}
	if len(nickname) == 0 {
		nickname = u.Nickname
	}
	if len(birthday) == 0 {
		birthday = u.Birthday
	}
	if len(details) == 0 {
		details = u.Details
		if len(details) == 0 {
			details = "此人很神秘，什么也没写"
		}
	}
	return r.dao.UpdateById(ctx, id, nickname, birthday, details)
}
