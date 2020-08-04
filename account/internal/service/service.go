package service

import (
	"context"
	"html/template"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/cli"
	"github.com/jinzhu/gorm"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/model"
)

type AccountService struct {
	mysqlCli *gorm.DB
	redisCli *redis.Client
	emailCli *cli.EmailCli

	captchaExpiration time.Duration
	accountExpiration time.Duration

	captchaEmailTpl *template.Template
}

func NewAccountService(mysqlCli *gorm.DB, redisCli *redis.Client, emailCli *cli.EmailCli, opts ...AccountServiceOption) (*AccountService, error) {
	options := defaultAccountServiceOptions
	for _, opt := range opts {
		opt(&options)
	}

	if !mysqlCli.HasTable(&model.Account{}) {
		if err := mysqlCli.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Account{}).Error; err != nil {
			return nil, err
		}
	}

	captchaEmailTpl, err := template.New("captcha").Parse(captchaTpl)
	if err != nil {
		return nil, err
	}

	return &AccountService{
		mysqlCli:        mysqlCli,
		redisCli:        redisCli,
		emailCli:        emailCli,
		captchaEmailTpl: captchaEmailTpl,

		captchaExpiration: options.CaptchaExpiration,
		accountExpiration: options.AccountExpiration,
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

type AccountServiceOptions struct {
	CaptchaExpiration time.Duration
	AccountExpiration time.Duration
}

var defaultAccountServiceOptions = AccountServiceOptions{
	CaptchaExpiration: 5 * time.Minute,
	AccountExpiration: 30 * time.Minute,
}

type AccountServiceOption func(options *AccountServiceOptions)

func WithCaptchaExpiration(captchaExpiration time.Duration) AccountServiceOption {
	return func(options *AccountServiceOptions) {
		options.CaptchaExpiration = captchaExpiration
	}
}

func WithAccountExpiration(accountExpiration time.Duration) AccountServiceOption {
	return func(options *AccountServiceOptions) {
		options.AccountExpiration = accountExpiration
	}
}
