package main

// This example launches an IPFS-Lite peer and fetches a hello-world
// hash from the IPFS network.repo

import (
	"context"
	"fmt"
	"github.com/topport/magic/internal/redcon"
	//crdt "github.com/ipfs/go-ds-crdt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"

	ipfslite "github.com/topport/magic"
	"github.com/topport/magic/internal/repo"
)
var addr=":6381"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//home, err := os.UserHomeDir()
	//if err != nil {
	//	return
	//}
	root := filepath.Join("./node")
	err := repo.Init(root, "4502")
	if err != nil {
		return
	}

	r, err := repo.Open(root)
	if err != nil {
		return
	}

	lite, err := ipfslite.New(ctx, cancel, r)
	if err != nil {
		panic(err)
	}
	fmt.Println(lite.Host.ID())
	addrs := []string{}
	for _, v := range lite.Host.Addrs() {
		if !strings.HasPrefix(v.String(), "127") {
			addrs = append(addrs, v.String()+"/p2p/"+lite.Host.ID().String())
		}
	}

	go redcon.Redconn(addr,lite)

	endWaiter := sync.WaitGroup{}
	endWaiter.Add(1)
	var sigChan chan os.Signal
	sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println()
		endWaiter.Done()
	}()
	endWaiter.Wait()
}
