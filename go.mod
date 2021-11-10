module github.com/topport/magic

go 1.15

require (
	github.com/asabya/go-ipc-uds v0.1.1
	github.com/bits-and-blooms/bloom/v3 v3.0.1
	github.com/brynbellomy/go-structomancer v1.0.3
	github.com/brynbellomy/klog v0.0.0-20200414031930-87fbf2e555ae
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/dgraph-io/badger/v2 v2.2007.4
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ethereum/go-ethereum v1.10.12
	github.com/facebookgo/atomicfile v0.0.0-20151019160806-2de1f203e7d5
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/grandcat/zeroconf v1.0.0
	github.com/h2non/filetype v1.1.1
	github.com/ipfs/go-bitswap v0.4.0
	github.com/ipfs/go-blockservice v0.1.7
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-datastore v0.4.6
	github.com/ipfs/go-ds-crdt v0.1.20
	github.com/ipfs/go-ds-leveldb v0.4.2
	github.com/ipfs/go-fs-lock v0.0.6
	github.com/ipfs/go-ipfs-blockstore v1.0.4
	github.com/ipfs/go-ipfs-chunker v0.0.5
	github.com/ipfs/go-ipfs-files v0.0.8 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.5
	github.com/ipfs/go-ipld-format v0.2.0
	github.com/ipfs/go-ipns v0.1.2
	github.com/ipfs/go-log/v2 v2.3.0
	github.com/ipfs/go-merkledag v0.4.0
	github.com/ipfs/go-unixfs v0.2.6
	github.com/ipfs/go-verifcid v0.0.1
	github.com/jbenet/goprocess v0.1.4
	github.com/libp2p/go-libp2p v0.15.1
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.9.0
	github.com/libp2p/go-libp2p-kad-dht v0.13.1
	github.com/libp2p/go-libp2p-pubsub v0.5.5
	github.com/libp2p/go-libp2p-record v0.1.3
	github.com/libp2p/go-libp2p-swarm v0.5.3
	github.com/libp2p/go-libp2p-tls v0.2.0
	github.com/libp2p/go-tcp-transport v0.2.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.4.1
	github.com/multiformats/go-multihash v0.0.16
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/plexsysio/taskmanager v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.8.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/redcon v1.4.2
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/urfave/cli v1.22.1
	go.uber.org/zap v1.19.0
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e
	golang.org/x/net v0.0.0-20211013171255-e13a2654a71e // indirect
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
	golang.org/x/tools v0.1.2 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/plexsysio/taskmanager => github.com/asabya/taskmanager v0.1.0
