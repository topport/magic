package command

import (
	"github.com/spf13/viper"
	"github.com/topport/magic/pkg/cli"
)

func newViper() *viper.Viper {
	return cli.NewViper("PYROSCOPE")
}
