package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topport/magic/cli/common"

	"github.com/topport/magic/pkg/cli"
	"github.com/topport/magic/pkg/config"
)

func newServerCmd(comm *common.Common,cfg *config.Server) *cobra.Command {
	vpr := newViper()
	serverCmd := &cobra.Command{
		Use:   "server [flags]",
		Short: "Start magic server. This is the database + web-based user interface",

		DisableFlagParsing: true,
		RunE: cli.CreateCmdRunFn(cfg, vpr, func(_ *cobra.Command, _ []string) error {

			fmt.Println(cfg.RedisBindAddr)
			return cli.StartServer(comm,cfg)
		}),
	}

	cli.PopulateFlagSet(cfg, serverCmd.Flags(), vpr)
	_ = serverCmd.Flags().MarkHidden("metrics-export-rules")
	return serverCmd
}
