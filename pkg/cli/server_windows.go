package cli

import (
	"fmt"

	"github.com/topport/magic/pkg/config"
)

func StartServer(_ *config.Server) error {
	return fmt.Errorf("server mode is not supported on Windows")
}
