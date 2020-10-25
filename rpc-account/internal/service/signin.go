package service

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/go-kit/strex"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-account/internal/model"
	"github.com/hatlonely/go-rpc/rpc-account/pkg/rpcx"
)

func GenerateToken() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}

func (s *AccountService) SignIn(ctx context.Context, req *account.SignInReq) (*account.SignInRes, error) {
	requestID := rpcx.MetaDataGetRequestID(ctx)

	a := &model.Account{}
	if strex.RePhone.MatchString(req.Username) {
		if err := s.mysqlCli.Where("phone=?", req.Username).First(a).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, rpcx.NewErrorf(codes.PermissionDenied, requestID, "Forbidden", "user [%v] not exist", req.Username)
			}
			return nil, errors.Wrapf(err, "mysql select user [%v] failed", req.Username)
		}
	} else if strex.ReEmail.MatchString(req.Username) {
		if err := s.mysqlCli.Where("email=?", req.Username).First(a).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, rpcx.NewErrorf(codes.PermissionDenied, requestID, "Forbidden", "user [%v] not exist", req.Username)
			}
			return nil, errors.Wrapf(err, "mysql select user [%v] failed", req.Username)
		}
	} else {
		return nil, rpcx.NewErrorf(codes.InvalidArgument, requestID, "InvalidArgument", "user [%v] is invalid", req.Username)
	}

	if a.Password != req.Password {
		return nil, rpcx.NewErrorf(codes.PermissionDenied, requestID, "Forbidden", "password is incorrect")
	}

	token := GenerateToken()

	if err := s.kv.Set(token, a); err != nil {
		return nil, errors.Wrapf(err, "redis set token [%v] failed", token)
	}

	return &account.SignInRes{
		Token: token,
	}, nil
}
