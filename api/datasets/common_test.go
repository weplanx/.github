package datasets_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"server/api"
	"server/bootstrap"
	"server/common"
	"testing"
	"time"
)

var (
	x *api.API
)

func TestMain(m *testing.M) {
	os.Chdir("../../")
	var err error
	var values *common.Values
	if values, err = bootstrap.LoadStaticValues("./config/default.values.yml"); err != nil {
		panic(err)
	}
	x, err = bootstrap.NewAPI(values)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if _, err = x.Initialize(ctx); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestService_Lists(t *testing.T) {
	ctx := context.TODO()
	data, err := x.DatasetsX.Lists(ctx, "")
	assert.NoError(t, err)
	t.Log(data)
}
