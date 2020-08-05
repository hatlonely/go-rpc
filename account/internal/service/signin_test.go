package service

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	account "github.com/hatlonely/go-rpc/account/api/gen/go/api"
)

func TestAccountService_SignIn(t *testing.T) {
	Convey("TestAccountService_SignIn", t, func() {
		res, err := service.SignIn(context.Background(), &account.SignInReq{
			Username: "hatlonely@fommail.com",
			Password: "12345678",
		})
		So(err, ShouldBeNil)
		So(len(res.Token), ShouldEqual, 32)
	})
}
