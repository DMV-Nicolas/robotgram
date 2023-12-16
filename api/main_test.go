package api

import (
	"testing"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, querier db.Querier) *Server {
	config := util.Config{}

	server, err := NewServer(config, querier)
	require.NoError(t, err)

	return server
}
