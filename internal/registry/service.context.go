package registry

import (
	"github.com/online_marketplace/helper/database"
	tokenprovider "github.com/online_marketplace/helper/token_provider"
	"github.com/online_marketplace/helper/token_provider/jwt"
	"github.com/online_marketplace/internal/config"
	"github.com/online_marketplace/internal/repository"
	"github.com/online_marketplace/internal/types/entity"
)

type ServiceContext struct {
	Config      config.Config
	UserRepo    repository.UserRepository
	ProductRepo repository.ProductRepository
	JwtProvider tokenprovider.Provider
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := database.Must(
		&c.Database,
		database.WithGormMigrate(entity.RegisterMigrate()),
	)
	userRepo := repository.NewUserRepository(sqlConn)
	productRepo := repository.NewProductRepository(sqlConn)
	return &ServiceContext{
		Config:      c,
		UserRepo:    userRepo,
		ProductRepo: productRepo,
		JwtProvider: jwt.NewTokenJWTProvider(c.JWT),
	}

}
