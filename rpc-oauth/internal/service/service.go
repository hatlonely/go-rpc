package service

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/strx"

	"github.com/hatlonely/go-rpc/rpc-oauth/api/gen/go/api"
)

type OAuthService struct {
}

func NewOAuthService() (*OAuthService, error) {
	return &OAuthService{}, nil
}

func (s *OAuthService) Auth(ctx context.Context, req *api.AuthReq) (*empty.Empty, error) {
	fmt.Println(strx.JsonMarshal(req))
	return &empty.Empty{}, nil
}

func (s *OAuthService) Token(ctx context.Context, req *api.TokenReq) (*api.TokenRes, error) {
	fmt.Println(strx.JsonMarshal(req))
	return &api.TokenRes{}, nil
}

func (s *OAuthService) CreateApp(ctx context.Context, req *api.CreateAppReq) (*empty.Empty, error) {
	fmt.Println(strx.JsonMarshal(req))
	return &empty.Empty{}, nil
}
