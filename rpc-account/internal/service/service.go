package service

import (
	"html/template"
	"time"

	"github.com/go-redis/redis"
	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/kv"
	"github.com/jinzhu/gorm"

	"github.com/hatlonely/go-rpc/rpc-account/internal/model"
)

type AccountService struct {
	mysqlCli *gorm.DB
	redisCli *redis.Client
	emailCli *cli.EmailCli
	kv       *kv.KV

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

	store, err := kv.NewRedisStore(redisCli, kv.WithRedisExpiration(options.AccountExpiration))
	if err != nil {
		return nil, err
	}
	kv := kv.NewKV(store, kv.WithStringKey())

	return &AccountService{
		mysqlCli:        mysqlCli,
		redisCli:        redisCli,
		emailCli:        emailCli,
		captchaEmailTpl: captchaEmailTpl,
		kv:              kv,

		captchaExpiration: options.CaptchaExpiration,
		accountExpiration: options.AccountExpiration,
	}, nil
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
