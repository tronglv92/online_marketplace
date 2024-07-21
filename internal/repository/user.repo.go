package repository

import (
	"context"

	"github.com/online_marketplace/helper/database"
	baseRepo "github.com/online_marketplace/helper/database/sqldata/repository"
	"github.com/online_marketplace/internal/types/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	First(ctx context.Context, opts ...baseRepo.QueryOpt) (*entity.User, error)
	WithEmail(email string) baseRepo.QueryOpt
	WithUID(uid string) baseRepo.QueryOpt
	CreateWithReturn(ctx context.Context, entity *entity.User) (*entity.User, error)
}

type userRepo struct {
	baseRepo.BaseRepository[entity.User]
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepo{
		baseRepo.NewBaseRepository[entity.User](db),
	}
}

func (ur *userRepo) WithEmail(email string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email=?", email)
	}
}

func (ur *userRepo) WithUID(uid string) baseRepo.QueryOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uid=?", uid)
	}
}
