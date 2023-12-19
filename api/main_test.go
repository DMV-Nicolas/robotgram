package api

import (
	"testing"
	"time"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, queries db.Querier, tokenSymmetricKey string) *Server {
	config := util.Config{
		TokenSymmetricKey:    tokenSymmetricKey,
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Minute * 2,
	}

	server, err := NewServer(config, queries)
	require.NoError(t, err)

	return server
}
