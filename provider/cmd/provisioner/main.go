package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/spf13/cobra"
	p "github.com/unmango/pulumi-baremetal/provider/pkg/provisioner"
)

var (
	address      string
	network      string
	clientCaFile string
	certFile     string
	keyFile      string
	verbose      bool
)

var rootCmd = &cobra.Command{
	Use:   "provisioner",
	Short: "The pulumi-baremetal provisioner",
	RunE: func(cmd *cobra.Command, args []string) error {
		var level slog.Level
		if verbose {
			level = slog.LevelDebug
		}

		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})

		log := slog.New(handler)
		lis, err := net.Listen(network, address)
		if err != nil {
			return fmt.Errorf("failed to create listener: %w", err)
		}

		log.Debug("creating provisioner")
		provisioner := p.New(lis,
			p.WithLogger(log),
			p.WithOptionalCertificates(clientCaFile, certFile, keyFile),
		)

		log.Info("serving",
			"network", network,
			"address", address,
			"verbose", verbose,
			"clientCaFile", clientCaFile,
			"certFile", certFile,
			"keyFile", keyFile,
		)

		return provisioner.Serve()
	},
}

func main() {
	rootCmd.Flags().StringVar(&address, "address", "", "Must be a valid `net.Listen()` address")
	rootCmd.Flags().StringVar(&network, "network", "tcp", "Must be a valid `net.Listen()` network. i.e. \"tcp\", \"tcp4\", \"tcp6\", \"unix\" or \"unixpacket\"")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Log verbosity")

	rootCmd.Flags().StringVar(&clientCaFile, "client-ca-file", "", "The path to the certificate authority file")
	rootCmd.Flags().StringVar(&certFile, "cert-file", "", "The path to the server certificate file")
	rootCmd.Flags().StringVar(&keyFile, "key-file", "", "The path to the server private key file")
	rootCmd.MarkFlagsRequiredTogether("client-ca-file", "cert-file", "key-file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("failed to execute: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("exiting gracefully")
}
