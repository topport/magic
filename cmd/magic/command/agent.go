package command

import (
	"github.com/spf13/cobra"

	"github.com/topport/magic/pkg/cli"
	"github.com/topport/magic/pkg/config"
)

func newAgentCmd(cfg *config.Agent) *cobra.Command {
	vpr := newViper()
	agentCmd := &cobra.Command{
		Use:   "agent [flags]",
		Short: "Start pyroscope agent",

		DisableFlagParsing: true,
		RunE: cli.CreateCmdRunFn(cfg, vpr, func(_ *cobra.Command, _ []string) error {
			//return cli.StartAgent(cfg)
			return nil
		}),
	}

	cli.PopulateFlagSet(cfg, agentCmd.Flags(), vpr, cli.WithSkip("targets"))
	return agentCmd
}
