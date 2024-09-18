package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"server/model"
	"testing"
)

func TestSetPictures(t *testing.T) {
	ctx := context.TODO()
	err := model.SetPictures(ctx, x.Db)
	assert.NoError(t, err)
}
