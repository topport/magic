package cli

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	ipfslite "github.com/topport/magic"
	"github.com/topport/magic/cli/common"
	"github.com/topport/magic/internal/redcon"
	"github.com/topport/magic/internal/repo"
	"github.com/topport/magic/pkg/config"
	"golang.org/x/sync/errgroup"
	"strings"

	"github.com/topport/magic/pkg/util/debug"
)

type serverService struct {
	config           *config.Server
	logger           *logrus.Logger
	//controller       *server.Controller
	ipfslite          *ipfslite.Peer
	//directUpstream   *direct.Direct
	//analyticsService *analytics.Service
	//selfProfiling    *agent.ProfileSession
	debugReporter    *debug.Reporter
//	healthController *health.Controller

	stopped chan struct{}
	done    chan struct{}
	group   *errgroup.Group
}

func newServerService(comm *common.Common,logger *logrus.Logger, c *config.Server) (*serverService, error) {
	svc := serverService{
		config:  c,
		logger:  logger,
		stopped: make(chan struct{}),
		done:    make(chan struct{}),
	}

	var err error
	fmt.Println( svc.config.APIBindAddr,svc.config.RedisBindAddr,svc.config.StoragePath,"asdfasfds")
	root:=svc.config.StoragePath
	err = repo.Init(root, svc.config.APIBindAddr)
	if err != nil {
		return nil, fmt.Errorf("new storage: %w", err)
	}
	//
	r, err := repo.Open(root)
	if err != nil {
		return nil, fmt.Errorf("new storage: %w", err)
	}
	fmt.Println(r)
	ctx, cancel := context.WithCancel(context.Background())
	svc.ipfslite, err = ipfslite.New(ctx, cancel, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(svc.ipfslite.Host.ID())
	addrs := []string{}
	for _, v := range svc.ipfslite.Host.Addrs() {
		if !strings.HasPrefix(v.String(), "127") {
			addrs = append(addrs, v.String()+"/p2p/"+svc.ipfslite.Host.ID().String())
		}
	}

	fmt.Println(addrs)
	go redcon.Redconn(svc.config.RedisBindAddr,svc.ipfslite)
	//svc.storage, err = storage.New(svc.config, prometheus.DefaultRegisterer)
	//if err != nil {
	//	return nil, fmt.Errorf("new storage: %w", err)
	//}



	//svc.healthController = health.NewController(svc.logger, time.Minute, diskPressure)
	//svc.debugReporter = debug.NewReporter(svc.logger, svc.storage, svc.config, prometheus.DefaultRegisterer)



	return &svc, nil
}

func (svc *serverService) Start() error {
	g, ctx := errgroup.WithContext(context.Background())
	svc.group = g


	//go svc.debugReporter.Start()


	svc.logger.Debug("collecting local profiles")
	//if err := svc.storage.CollectLocalProfiles(); err != nil {
	//	svc.logger.WithError(err).Error("failed to collect local profiles")
	//}

	defer close(svc.done)
	select {
	case <-svc.stopped:
	case <-ctx.Done():
		// The context is canceled the first time a function passed to Go
		// returns a non-nil error.
	}
	// N.B. internal components are de-initialized/disposed (if applicable)
	// regardless of the exit reason. Once server is stopped, wait for all
	// Go goroutines to finish.
	svc.stop()
	return svc.group.Wait()
}

func (svc *serverService) Stop() {
	close(svc.stopped)
	<-svc.done
}

//revive:disable-next-line:confusing-naming methods are different
func (svc *serverService) stop() {

	//svc.debugReporter.Stop()

	svc.logger.Debug("stopping storage")
	//if err := svc.storage.Close(); err != nil {
	//	svc.logger.WithError(err).Error("storage close")
	//}
	svc.logger.Debug("stopping http server")

}
