package common

import (
	"context"

	ipfslite "github.com/topport/magic"
	"github.com/topport/magic/internal/repo"
)

// Common features for cli commands
type Common struct {
	Root     string
	Repo     repo.Repo
	LitePeer *ipfslite.Peer
	Context  context.Context
	Cancel   context.CancelFunc
}
