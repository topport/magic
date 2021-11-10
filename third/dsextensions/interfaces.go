package dsextensions

import (
	"context"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

type QueryExt struct {
	query.Query
	SeekPrefix string
}

type TxnExt interface {
	datastore.Datastore
	datastore.Txn
	QueryExtensions
}



type DatastoreExtensions interface {
	datastore.Datastore

	NewTransactionExtended(ctx context.Context, readOnly bool) (TxnExt, error)
	QueryExtensions
}

type QueryExtensions interface {
	QueryExtended(ctx context.Context, q QueryExt) (query.Results, error)
}
