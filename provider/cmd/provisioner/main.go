package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/spf13/cobra"
	p "github.com/unmango/pulumi-baremetal/provider/pkg/provisioner"
)

var (
	address string
	network string
	verbose bool
	log     *slog.Logger
)

var rootCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "The pulumi-baremetal provisioner",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log = logger(verbose)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		lis, err := net.Listen(network, address)
		if err != nil {
			return err
		}

		log.Debug("creating provisioner")
		provisioner := p.New(lis, p.WithLogger(log))

		log.Info("serving", "network", network, "address", address, "verbose", verbose)
		return provisioner.Serve()
	},
}

func main() {
	rootCmd.Flags().StringVar(&address, "address", "", "Must be a valid `net.Listen()` address.")
	rootCmd.Flags().StringVar(&network, "network", "tcp", "Must be a valid `net.Listen()` network. i.e. \"tcp\", \"tcp4\", \"tcp6\", \"unix\" or \"unixpacket\"")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Log verbosity")

	if err := rootCmd.Execute(); err != nil {
		log.Error("failed to execute", "err", err)
	}

	log.Debug("exiting gracefully")
}

func logger(verbose bool) *slog.Logger {
	var level slog.Level
	if verbose {
		level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)
}
