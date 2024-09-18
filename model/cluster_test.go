package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"server/model"
	"testing"
)

func TestSetClusters(t *testing.T) {
	ctx := context.TODO()
	err := model.SetClusters(ctx, x.Db)
	assert.NoError(t, err)
}
