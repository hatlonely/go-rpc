package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/hatlonely/go-kit/strex"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/account/internal/model"
)

func GenerateToken() string {
	buf := make([]byte, 32)
	_, _ = rand.Read(buf)
	token := make([]byte, 16)
	hex.Encode(buf, token)
	return string(buf)
}

func (s *AccountService) SignIn(ctx context.Context, req *account.SignInReq) (*account.SignInRes, error) {
	a := &model.Account{}
	if strex.RePhone.MatchString(req.Username) {
		if err := s.mysqlCli.Where("phone=?", req.Username).First(a).Error; err != nil {
			return nil, err
		}
	} else if strex.ReEmail.MatchString(req.Username) {
		if err := s.mysqlCli.Where("email=?", req.Username).First(a).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid username [%v]", req.Username)
	}

	if a.Password != req.Password {
		return nil, status.Errorf(codes.PermissionDenied, "password not correct")
	}

	token := GenerateToken()

	if err := s.redisCli.Set(token, a, s.accountExpiration).Err(); err != nil {
		return nil, errors.Wrapf(err, "redis set token [%v] failed", token)
	}

	return &account.SignInRes{
		Token: token,
	}, nil
}
