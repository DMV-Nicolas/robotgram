package api

import (
	"testing"
	"time"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, querier db.Querier) *Server {
	config := util.Config{
		TokenSymmetricKey:    util.RandomPassword(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Minute * 2,
	}

	server, err := NewServer(config, querier)
	require.NoError(t, err)

	return server
}
