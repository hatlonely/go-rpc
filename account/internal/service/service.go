package service

import (
	"context"
	"html/template"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/cli"
	"github.com/jinzhu/gorm"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/model"
)

type AccountService struct {
	mysql *gorm.DB
	redis *redis.Client
	email *cli.EmailCli

	captchaEmailTpl *template.Template
}

func NewAccountService(mysql *gorm.DB, redis *redis.Client, email *cli.EmailCli) (*AccountService, error) {
	if !mysql.HasTable(&model.Account{}) {
		if err := mysql.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Account{}).Error; err != nil {
			return nil, err
		}
	}

	captchaEmailTpl, err := template.New("captcha").Parse(captchaTpl)
	if err != nil {
		return nil, err
	}

	return &AccountService{
		mysql:           mysql,
		redis:           redis,
		email:           email,
		captchaEmailTpl: captchaEmailTpl,
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
