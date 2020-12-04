package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExec(t *testing.T) {
	Convey("TestExec", t, func() {
		Convey("case shell", func() {
			exitCode, stdout, stderr, err := Exec("sh", "echo hello world")
			So(exitCode, ShouldEqual, 0)
			So(stdout, ShouldEqual, "hello world\n")
			So(stderr, ShouldBeEmpty)
			So(err, ShouldBeNil)
		})

		Convey("case python", func() {
			exitCode, stdout, stderr, err := Exec("python3", "print('hello world')")
			So(exitCode, ShouldEqual, 0)
			So(stdout, ShouldEqual, "hello world\n")
			So(stderr, ShouldBeEmpty)
			So(err, ShouldBeNil)
		})
	})
}
