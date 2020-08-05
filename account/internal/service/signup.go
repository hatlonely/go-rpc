package service

import (
	"context"
	"strings"

	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/model"
)

func (s *AccountService) SignUp(ctx context.Context, req *account.SignUpReq) (*empty.Empty, error) {
	key := "captcha_" + req.Email
	val, err := s.redisCli.Get(key).Result()
	if err == redis.Nil {
		return nil, status.Errorf(codes.InvalidArgument, "captcha is not exists")
	}
	if err != nil {
		return nil, errors.Wrapf(err, "redis get key [%v] failed", key)
	}
	if req.Captcha != val {
		return nil, status.Errorf(codes.InvalidArgument, "captcha is not match")
	}

	birthday, err := cast.ToTimeE(req.Birthday)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid birthday format")
	}

	user := &model.Account{
		Email:    req.Email,
		Phone:    req.Phone,
		Name:     req.Name,
		Password: req.Password,
		Birthday: birthday,
		Gender:   int(req.Gender),
		Avatar:   req.Avatar,
	}
	if err := s.mysqlCli.Create(user).Error; err != nil {
		switch err.(type) {
		case *mysql.MySQLError:
			e := err.(*mysql.MySQLError)
			if e.Number == 1062 {
				if strings.Contains(e.Message, "accounts.email_idx") {
					return nil, status.Errorf(codes.InvalidArgument, "account [%v] exists", req.Email)
				}
				if strings.Contains(e.Message, "accounts.phone_idx") {
					return nil, status.Errorf(codes.InvalidArgument, "account [%v] exists", req.Phone)
				}
				if strings.Contains(e.Message, "accounts.name_idx") {
					return nil, status.Errorf(codes.InvalidArgument, "account [%v] exists", req.Name)
				}
			}
		}
		return nil, errors.Wrap(err, "mysql create failed")
	}

	return &empty.Empty{}, nil
}