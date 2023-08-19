package arbitrum

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// usc68059@omeie.com

func TestNew(t *testing.T) {
	client := New(&Config{})

	res, err := client.GetListTx(context.Background(),
		"ARBI_USDT",
		"0x09197c3dd57E86Cb8b02A7ca2c315a7e59dE9383",
		0,
	)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
