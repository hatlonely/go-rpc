package service

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/go-kit/strex"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-account/internal/model"
)

func GenerateToken() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}

func (s *AccountService) SignIn(ctx context.Context, req *account.SignInReq) (*account.SignInRes, error) {
	a := &model.Account{}
	if strex.RePhone.MatchString(req.Username) {
		if err := s.mysqlCli.Where("phone=?", req.Username).First(a).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, status.Errorf(codes.PermissionDenied, "user [%v] not exist", req.Username)
			}
			return nil, err
		}
	} else if strex.ReEmail.MatchString(req.Username) {
		if err := s.mysqlCli.Where("email=?", req.Username).First(a).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, status.Errorf(codes.PermissionDenied, "user [%v] not exist", req.Username)
			}
			return nil, err
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid username [%v]", req.Username)
	}

	if a.Password != req.Password {
		return nil, status.Errorf(codes.PermissionDenied, "password not correct")
	}

	token := GenerateToken()

	if err := s.kv.Set(token, a); err != nil {
		return nil, errors.Wrapf(err, "redis set token [%v] failed", token)
	}

	return &account.SignInRes{
		Token: token,
	}, nil
}
