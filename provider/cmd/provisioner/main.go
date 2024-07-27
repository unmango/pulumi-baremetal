package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner"
)

var log = logger()

var (
	address string
	network string
)

var rootCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "The pulumi-baremetal provisioner",
	RunE: func(cmd *cobra.Command, args []string) error {
		lis, err := net.Listen(network, address)
		if err != nil {
			return err
		}

		log.Info("Serving requests", "network", network, "address", address)
		return provisioner.Serve(lis)
	},
}

func main() {
	rootCmd.Flags().StringVar(&address, "address", "", "Must be a valid `net.Listen()` address.")
	rootCmd.Flags().StringVar(&network, "network", "tcp", "Must be a valid `net.Listen()` network. i.e. \"tcp\", \"tcp4\", \"tcp6\", \"unix\" or \"unixpacket\"")

	log.Debug("Starting provisioner")
	if err := rootCmd.Execute(); err != nil {
		log.Error("provisioner failed", "err", err)
	}
}

func logger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	return slog.New(handler)
}
