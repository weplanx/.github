package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"server/model"
	"testing"
)

func TestSetImessages(t *testing.T) {
	ctx := context.TODO()
	err := model.SetImessages(ctx, x.Db)
	assert.NoError(t, err)
}
