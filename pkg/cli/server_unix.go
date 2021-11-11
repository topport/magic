// +build !windows

package cli

import (
	"fmt"
	"github.com/topport/magic/cli/common"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/topport/magic/pkg/config"
)

func StartServer(comm *common.Common,c *config.Server) error {
	logLevel, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLevel)
	logger := logrus.StandardLogger()
	srv, err := newServerService(comm,logger, c)
	if err != nil {
		return fmt.Errorf("could not initialize server: %w", err)
	}



	var stopTime time.Time
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	exited := make(chan error)
	go func() {
		exited <- srv.Start()
		close(exited)
	}()

	select {
	case <-s:
		logger.Info("stopping server")
		stopTime = time.Now()
		srv.Stop()
		err = <-exited
		if err != nil {
			logger.WithError(err).Error("failed to stop server gracefully")
			return err
		}
		logger.WithField("duration", time.Since(stopTime)).Info("server stopped gracefully")
		return nil

	case err = <-exited:
		if err == nil {
			// Should never happen.
			logger.Error("server exited")
			return nil
		}
		logger.WithError(err).Error("server failed")
		return err
	}
}
