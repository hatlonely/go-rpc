package service

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/model"
)

type AccountService struct {
	mysql *gorm.DB
	redis *redis.Client
}

func NewAccountService(mysql *gorm.DB, redis *redis.Client) (*AccountService, error) {
	if !mysql.HasTable(&model.Account{}) {
		if err := mysql.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Account{}).Error; err != nil {
			return nil, err
		}
	}

	return &AccountService{
		mysql: mysql,
		redis: redis,
	}, nil
}

func (s *AccountService) SignIn(ctx context.Context, req *account.SignInReq) (*account.SignInRes, error) {
	return &account.SignInRes{
		Token: "hello world",
	}, nil
}

func (s *AccountService) SignUp(ctx context.Context, req *account.SignUpReq) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}

func (s *AccountService) SignOut(ctx context.Context, req *account.SignOutReq) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
