package main

import (
	"os"

	opservice "github.com/tokamak-network/tokamak-thanos/op-service"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
	"github.com/tokamak-network/tokamak-thanos/op-proposer/flags"
	"github.com/tokamak-network/tokamak-thanos/op-proposer/metrics"
	"github.com/tokamak-network/tokamak-thanos/op-proposer/proposer"
	"github.com/tokamak-network/tokamak-thanos/op-service/cliapp"
	oplog "github.com/tokamak-network/tokamak-thanos/op-service/log"
	"github.com/tokamak-network/tokamak-thanos/op-service/metrics/doc"
)

var (
	Version   = "v0.10.14"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "op-proposer"
	app.Usage = "L2 Output Submitter"
	app.Description = "Service for generating and proposing L2 Outputs"
	app.Action = cliapp.LifecycleCmd(proposer.Main(Version))
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.NewSubcommands(metrics.NewMetrics("default")),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
