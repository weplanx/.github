package model_test

import (
	"github.com/weplanx/server/api"
	"github.com/weplanx/server/bootstrap"
	"os"
	"testing"
)

var x *api.API

func TestMain(m *testing.M) {
	os.Chdir("../")
	values, err := bootstrap.LoadStaticValues("./config/default.values.yml")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	if x, err = bootstrap.NewAPI(values); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
